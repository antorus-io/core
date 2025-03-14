package storage

import "encoding/json"

type StorageType string

const (
	Redis StorageType = "REDIS"
)

var StorageInitialized = false

type StorageClient interface {
	Del(namespace string, key string) error
	Get(key string) (*string, error)
	GetAllFromNamespace(namespace string) (map[string]json.RawMessage, error)
	Ping() error
	Publish(channel string, payload interface{}) error
	Set(namespace string, id string, value any) error
	Subscribe(channel string, handler func(string)) error
}
