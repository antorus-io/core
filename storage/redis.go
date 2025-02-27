package storage

import (
	"context"
	"fmt"

	"github.com/antorus-io/core/config"
	"github.com/antorus-io/core/logs"
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

func (r *RedisStorage) Set(namespace string, id string, value string) error {
	err := r.client.Set(r.ctx, fmt.Sprintf("%s:%s", namespace, id), value, 0).Err()

	return err
}

func (r *RedisStorage) Subscribe(channel string, handler func(payload string)) error {
	pubsub := r.client.Subscribe(r.ctx, channel)

	logs.Logger.Info("Subscribed to Redis channel", "channel", channel)

	go func() {
		for {
			msg, err := pubsub.ReceiveMessage(r.ctx)

			if err != nil {
				logs.Logger.Error("Error receiving Redis message", "error", err)

				continue
			}

			handler(msg.Payload)
		}
	}()

	return nil
}
