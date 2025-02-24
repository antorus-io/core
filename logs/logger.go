package logs

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/antorus-io/core/config"
	"github.com/antorus-io/core/models"
)

type LogLevel string

const (
	Debug LogLevel = "INFO"
	Error LogLevel = "ERROR"
	Info  LogLevel = "INFO"
	Warn  LogLevel = "WARN"
)

var Logger *LogHandler

type LogHandler struct {
	app    *config.ApplicationConfig
	logger *slog.Logger
}

func CreateLogger(app *config.ApplicationConfig) {
	Logger = &LogHandler{app, slog.New(slog.NewJSONHandler(os.Stdout, nil))}
}

func (l *LogHandler) Debug(msg string, params ...any) {
	baseParams := []any{"applicationMode", string(l.app.Mode), "env", l.app.Env, "service", l.app.Service}
	allParams := append(baseParams, params...)

	l.logger.Debug(msg, allParams...)

	if l.app.InitConfig.Database {
		l.saveToDatabase(Debug, msg, params...)
	}
}

func (l *LogHandler) Error(msg string, params ...any) {
	baseParams := []any{"applicationMode", string(l.app.Mode), "env", l.app.Env, "service", l.app.Service}
	allParams := append(baseParams, params...)

	l.logger.Error(msg, allParams...)

	if l.app.InitConfig.Database {
		l.saveToDatabase(Error, msg, params...)
	}
}

func (l *LogHandler) Info(msg string, params ...any) {
	baseParams := []any{"applicationMode", string(l.app.Mode), "env", l.app.Env, "service", l.app.Service}
	allParams := append(baseParams, params...)

	l.logger.Info(msg, allParams...)

	if l.app.InitConfig.Database {
		l.saveToDatabase(Info, msg, params...)
	}
}

func (l *LogHandler) Warn(msg string, params ...any) {
	baseParams := []any{"applicationMode", string(l.app.Mode), "env", l.app.Env, "service", l.app.Service}
	allParams := append(baseParams, params...)

	l.logger.Warn(msg, allParams...)

	if l.app.InitConfig.Database {
		l.saveToDatabase(Warn, msg, params...)
	}
}

func (l *LogHandler) saveToDatabase(level LogLevel, msg string, ps ...any) {
	params, err := parseParams(ps...)

	if err != nil {
		l.logger.Error("Failed to parse params", "error", err.Error())

		params = "{}"
	}

	logEntry := &models.LogEntry{
		ApplicationMode: string(l.app.Mode),
		Env:             l.app.Env,
		Level:           string(level),
		Message:         msg,
		Params:          params,
		Service:         l.app.Service,
		Timestamp:       time.Now(),
	}
	err = l.app.Models.LogEntries.Insert(logEntry)

	if err != nil {
		l.logger.Error("An error occurred during database operation", "error", err.Error())
	}
}

func parseParams(params ...any) (string, error) {
	if len(params)%2 != 0 {
		return "", fmt.Errorf("params must be in key-value pairs")
	}

	ps := make(map[string]interface{})

	for i := 0; i < len(params); i += 2 {
		key, ok := params[i].(string)

		if !ok {
			return "", fmt.Errorf("key must be a string")
		}

		ps[key] = params[i+1]
	}

	paramsJSON, err := json.Marshal(params)

	if err != nil {
		return "", err
	}

	return string(paramsJSON), nil
}
