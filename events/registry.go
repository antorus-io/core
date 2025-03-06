package events

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"maps"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Event struct {
	CreatedAt      time.Time
	Description    string
	Key            string
	SourceService  string
	TargetServices []string
	UpdatedAt      time.Time
	Version        int
}

type EventRegistry struct {
	events map[string]Event
	mu     sync.RWMutex
}

var (
	EventRegistryInitialized bool = false
	registryInstance         *EventRegistry
	once                     sync.Once
)

func GetEventRegistry() *EventRegistry {
	if registryInstance == nil {
		panic("EventRegistry not initialized. Call InitEventRegistry first.")
	}

	return registryInstance
}

func InitEventRegistry(ctx context.Context, pool *pgxpool.Pool) error {
	var initErr error

	once.Do(func() {
		registry := &EventRegistry{
			events: make(map[string]Event),
		}

		events, err := loadEventsFromDB(ctx, pool)

		if err != nil {
			initErr = fmt.Errorf("failed to load events: %w", err)

			return
		}

		for _, e := range events {
			registry.events[e.Key] = e
		}

		registryInstance = registry
		EventRegistryInitialized = true
	})

	return initErr
}

func (r *EventRegistry) AllEvents() map[string]Event {
	r.mu.RLock()
	defer r.mu.RUnlock()

	eventsMap := make(map[string]Event, len(r.events))

	maps.Copy(eventsMap, r.events)

	return eventsMap
}

func (r *EventRegistry) GetEvent(key string) (Event, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	event, exists := r.events[key]

	return event, exists
}

func loadEventsFromDB(ctx context.Context, pool *pgxpool.Pool) ([]Event, error) {
	rows, err := pool.Query(ctx, `
		SELECT created_at, description, event_key, source_service, target_services, updated_at, version
		FROM event_registry
	`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []Event

	for rows.Next() {
		var e Event
		var targets string

		err := rows.Scan(&e.CreatedAt, &e.Description, &e.Key, &e.SourceService, &targets, &e.UpdatedAt, &e.Version)

		if err != nil {
			return nil, err
		}

		e.TargetServices = splitCSV(targets)
		events = append(events, e)
	}

	return events, rows.Err()
}

func splitCSV(csv string) []string {
	if csv == "" {
		return nil
	}

	return strings.Split(csv, ",")
}
