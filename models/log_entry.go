package models

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type LogEntryModel struct {
	Pool *pgxpool.Pool
}

type LogEntry struct {
	ApplicationMode string    `json:"applicationMode" db:"application_mode"`
	Env             string    `json:"env" db:"env"`
	Id              string    `json:"id" db:"id"`
	Level           string    `json:"level" db:"level"`
	Message         string    `json:"message" db:"message"`
	Params          string    `json:"params" db:"params"`
	Service         string    `json:"service" db:"service"`
	Timestamp       time.Time `json:"timestamp" db:"timestamp"`
}

func (m LogEntryModel) Insert(l *LogEntry) error {
	query := `
		INSERT INTO log_entries (application_mode, env, level, message, params, service, timestamp)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id
	`
	args := []interface{}{l.ApplicationMode, l.Env, l.Level, l.Message, l.Params, l.Service, l.Timestamp}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	err := m.Pool.QueryRow(ctx, query, args...).Scan(&l.Id)

	if err != nil {
		return err
	}

	return nil
}
