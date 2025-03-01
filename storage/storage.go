package storage

type StorageType string

const (
	Redis StorageType = "REDIS"
)

type StorageClient interface {
	Get(key string) (*string, error)
	Ping() error
	Publish(channel string, payload interface{}) error
	Set(namespace string, id string, value string) error
	Subscribe(channel string, handler func(string)) error
}
