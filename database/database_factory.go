package database

import (
	"fmt"
	"strings"

	"github.com/antorus-io/core/config"
)

func CreateDatabase(cfg config.DatabaseConfig) error {
	if strings.ToUpper(cfg.Driver) == Postgres {
		err := CreatePostgresDatabase(cfg)

		if err != nil {
			return err
		}

		return nil
	}

	return fmt.Errorf("unsupported database driver: %s", cfg.Driver)
}
