package storage

import (
	"context"
	"encoding/json"
	"fmt"

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
		return fmt.Errorf("Failed to publish message", "channel", channel, "error", err)
	}

	return nil
}

func (r *RedisStorage) Set(namespace string, id string, value string) error {
	err := r.client.Set(r.ctx, fmt.Sprintf("%s:%s", namespace, id), value, 0).Err()

	return err
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
