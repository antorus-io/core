package database

import (
	"context"
	"fmt"
	"time"

	"github.com/antorus-io/core/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

var DatabaseInstance Database

type PostgresDatabase struct {
	Config     config.DatabaseConfig
	Connection *pgxpool.Pool
}

func CreatePostgresDatabase(cfg config.DatabaseConfig) error {
	db := &PostgresDatabase{Config: cfg}
	pool, err := db.openDB()

	if err != nil {
		return fmt.Errorf("failed to initialize Postgres database: %w", err)
	}

	db.Connection = pool
	DatabaseInstance = db
	DatabaseInitialized = true

	return nil
}

func (db *PostgresDatabase) Close() {
	if db.Connection != nil {
		db.Connection.Close()
	}
}

func (db *PostgresDatabase) GetPool() *pgxpool.Pool {
	return db.Connection
}

func (db *PostgresDatabase) OpenDB() (*pgxpool.Pool, error) {
	pool, err := db.openDB()

	if err != nil {
		return nil, err
	}

	db.Connection = pool

	return pool, nil
}

func (db *PostgresDatabase) openDB() (*pgxpool.Pool, error) {
	dsn := formatDsn(db.Config)
	config, err := pgxpool.ParseConfig(dsn)

	if err != nil {
		return nil, fmt.Errorf("unable to parse config: %w", err)
	}

	config.MaxConns = db.Config.MaxOpenConns
	maxIdleTime, err := time.ParseDuration(db.Config.MaxIdleTime)

	if err != nil {
		return nil, fmt.Errorf("invalid max idle time: %w", err)
	}

	config.MaxConnLifetime = maxIdleTime

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.NewWithConfig(ctx, config)

	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()

		return nil, fmt.Errorf("unable to ping database: %w", err)
	}

	return pool, nil
}

func formatDsn(db config.DatabaseConfig) string {
	return fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=%s", db.Driver, db.User, db.Password, db.Host, db.Port, db.Name, db.Sslmode)
}
