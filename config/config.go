package config

import (
	"log/slog"
	"os"
	"strconv"

	"github.com/antorus-io/core/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Provide a temporary logger as the proper one is not instantiated yet.
var tmpLogger *slog.Logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

type ApplicationMode string

const (
	Development ApplicationMode = "DEVELOPMENT"
	Production  ApplicationMode = "PRODUCTION"
)

type ApplicationConfig struct {
	DatabaseConfig DatabaseConfig
	Env            string
	InitConfig     CoreInitConfig
	Mode           ApplicationMode
	Models         models.Models
	Service        string
	StorageConfig  StorageConfig
}

type CoreInitConfig struct {
	Database bool
	Logger   bool
	Storage  bool
}

type DatabaseConfig struct {
	Driver       string
	Host         string
	MaxIdleConns int
	MaxIdleTime  string
	MaxOpenConns int32
	Name         string
	Password     string
	Port         string
	Sslmode      string
	User         string
}

type StorageConfig struct {
	Host string
	Port string
	Type string
}

func Setup(coreInitConfig CoreInitConfig) *ApplicationConfig {
	app := &ApplicationConfig{
		InitConfig: coreInitConfig,
	}

	if coreInitConfig.Database {
		app.setupDatabaseConfig()
	}

	app.setupApplicationEnvironment()

	if coreInitConfig.Storage {
		app.setupStorageConfig()
	}

	return app
}

func (app *ApplicationConfig) SetupModels(pool *pgxpool.Pool) {
	app.Models = models.NewModels(pool)
}

func (app *ApplicationConfig) setupApplicationEnvironment() {
	app.Env = os.Getenv("APPLICATION_ENV")
	app.Mode = Development
	app.Service = os.Getenv("SERVICE_NAME")

	if os.Getenv("APPLICATION_MODE") == string(Production) {
		app.Mode = Production
	}
}

func (app *ApplicationConfig) setupDatabaseConfig() {
	app.DatabaseConfig.Driver = "postgres"
	app.DatabaseConfig.Host = "postgres"
	app.DatabaseConfig.MaxIdleConns = 15
	app.DatabaseConfig.MaxIdleTime = "15m"
	app.DatabaseConfig.MaxOpenConns = 15
	app.DatabaseConfig.Name = "antorus"
	app.DatabaseConfig.Password = "pass1234"
	app.DatabaseConfig.Sslmode = "disable"
	app.DatabaseConfig.Port = "5432"
	app.DatabaseConfig.User = "antorus"

	if os.Getenv("DB_DRIVER") != "" {
		app.DatabaseConfig.Driver = os.Getenv("DB_DRIVER")
	}

	if os.Getenv("DB_HOST") != "" {
		app.DatabaseConfig.Host = os.Getenv("DB_HOST")
	}

	if os.Getenv("DB_NAME") != "" {
		app.DatabaseConfig.Name = os.Getenv("DB_NAME")
	}

	if os.Getenv("DB_PASSWORD") != "" {
		app.DatabaseConfig.Password = os.Getenv("DB_PASSWORD")
	}

	if os.Getenv("DB_PORT") != "" {
		app.DatabaseConfig.Port = os.Getenv("DB_PORT")
	}

	if os.Getenv("DB_SSLMODE") != "" {
		app.DatabaseConfig.Sslmode = os.Getenv("DB_SSLMODE")
	}

	if os.Getenv("DB_USER") != "" {
		app.DatabaseConfig.User = os.Getenv("DB_USER")
	}

	if os.Getenv("DB_MAX_IDLE_CONNS") != "" {
		maxIdleConns, err := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNS"))

		if err != nil {
			tmpLogger.Error(err.Error())
		}

		app.DatabaseConfig.MaxIdleConns = maxIdleConns
	}

	if os.Getenv("DB_MAX_IDLE_TIME") != "" {
		app.DatabaseConfig.MaxIdleTime = os.Getenv("DB_MAX_IDLE_TIME")
	}

	if os.Getenv("DB_MAX_OPEN_CONNS") != "" {
		maxOpenConns, err := strconv.ParseInt(os.Getenv("DB_MAX_OPEN_CONNS"), 10, 32)

		if err != nil {
			tmpLogger.Error(err.Error())

			maxOpenConns = 15
		}

		app.DatabaseConfig.MaxOpenConns = int32(maxOpenConns)
	}
}

func (app *ApplicationConfig) setupStorageConfig() {
	app.StorageConfig.Host = "0.0.0.0"
	app.StorageConfig.Port = "6379"
	app.StorageConfig.Type = "REDIS"

	if os.Getenv("STORAGE_HOST") != "" {
		app.StorageConfig.Host = os.Getenv("STORAGE_HOST")
	}

	if os.Getenv("STORAGE_PORT") != "" {
		app.StorageConfig.Port = os.Getenv("STORAGE_PORT")
	}

	if os.Getenv("STORAGE_TYPE") != "" {
		app.StorageConfig.Type = os.Getenv("STORAGE_TYPE")
	}
}
