package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/antorus-io/core/config"
	"github.com/redis/go-redis/v9"
)

var StorageInstance StorageClient

type RedisStorage struct {
	client *redis.Client
	ctx    context.Context
}

func CreateRedisStorage(storageConfig config.StorageConfig) error {
	client := &RedisStorage{
		client: redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", storageConfig.Host, storageConfig.Port),
			DB:       0,
			Password: "",
		}),
		ctx: context.Background(),
	}

	err := client.Ping()

	if err != nil {
		return err
	}

	StorageInstance = client
	StorageInitialized = true

	return nil
}

func (r *RedisStorage) Del(namespace string, key string) error {
	_, err := r.client.Del(r.ctx, fmt.Sprintf("%s:%s", namespace, key)).Result()

	if err != nil {
		return err
	}

	return nil
}

func (r *RedisStorage) Get(key string) (*string, error) {
	result, err := r.client.Get(r.ctx, key).Result()

	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}

		return nil, err
	}

	return &result, nil
}

func (r *RedisStorage) GetAllFromNamespace(namespace string) (map[string]json.RawMessage, error) {
	iter := r.client.Scan(r.ctx, 0, namespace+":*", 0).Iterator()
	result := make(map[string]json.RawMessage)

	for iter.Next(r.ctx) {
		fullKey := iter.Val()
		value, err := r.client.Get(r.ctx, fullKey).Result()

		if err != nil {
			return nil, err
		}

		keyWithoutNamespace := strings.TrimPrefix(fullKey, namespace+":")
		result[keyWithoutNamespace] = json.RawMessage(value)
	}

	if err := iter.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *RedisStorage) Ping() error {
	_, err := r.client.Ping(r.ctx).Result()

	return err
}

func (r *RedisStorage) Publish(channel string, payload interface{}) error {
	data, err := json.Marshal(payload)

	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	err = r.client.Publish(r.ctx, channel, string(data)).Err()

	if err != nil {
		return err
	}

	return nil
}

func (r *RedisStorage) Set(namespace string, id string, value any) error {
	fullKey := namespace + ":" + id

	switch v := value.(type) {
	case string:
		return r.client.Set(r.ctx, fullKey, v, 0).Err()
	default:
		data, err := json.Marshal(value)

		if err != nil {
			return fmt.Errorf("failed to marshal value: %w", err)
		}

		return r.client.Set(r.ctx, fullKey, data, 0).Err()
	}
}

func (r *RedisStorage) Subscribe(channel string, handler func(payload string)) error {
	pubsub := r.client.Subscribe(r.ctx, channel)

	go func() {
		for {
			msg, err := pubsub.ReceiveMessage(r.ctx)

			if err != nil {
				continue
			}

			handler(msg.Payload)
		}
	}()

	return nil
}
