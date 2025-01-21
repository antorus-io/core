package models

import "github.com/jackc/pgx/v5/pgxpool"

type Models struct {
	LogEntries LogEntryModel
}

func NewModels(pool *pgxpool.Pool) Models {
	return Models{
		LogEntries: LogEntryModel{Pool: pool},
	}
}
