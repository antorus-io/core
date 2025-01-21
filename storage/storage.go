package storage

type StorageType string

const (
	Redis StorageType = "REDIS"
)

type StorageClient interface {
	Get(key string) (*string, error)
	Ping() error
	Set(namespace string, id string, value string) error
}
