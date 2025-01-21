package storage

import (
	"fmt"
	"strings"

	"github.com/antorus-io/core/config"
)

func CreateStorage(cfg config.StorageConfig) error {
	if StorageType(strings.ToUpper(cfg.Type)) == Redis {
		err := CreateRedisStorage(cfg)

		if err != nil {
			return err
		}

		return nil
	}

	return fmt.Errorf("unsupported storage type: %s", cfg.Type)
}
