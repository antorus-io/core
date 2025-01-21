package database

import "github.com/jackc/pgx/v5/pgxpool"

const (
	Postgres = "POSTGRES"
)

// TODO: make arguments these DB-agnostic
type Database interface {
	Close()
	GetPool() *pgxpool.Pool
	OpenDB() (*pgxpool.Pool, error)
}
