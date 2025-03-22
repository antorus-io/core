package config

import (
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/antorus-io/core/events"
	"github.com/antorus-io/core/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Provide a temporary logger as the proper one is not instantiated yet.
var tmpLogger *slog.Logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

const WILDCARD_ADDR = "0.0.0.0"

type ApplicationConfig struct {
	DatabaseConfig DatabaseConfig
	Env            string
	Events         map[string]events.Event
	InitConfig     CoreInitConfig
	Models         models.Models
	ServerConfig   ServerConfig
	Service        string
	StorageConfig  StorageConfig
}

type CoreInitConfig struct {
	Database bool
	Logger   bool
	Server   bool
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

type Route struct {
	Handler http.HandlerFunc
	Path    string
}
type RouteConfig map[string]Route

type ServerConfig struct {
	Debug          string
	Host           string
	Port           string
	Routes         RouteConfig
	TrustedOrigins []string
}

type StorageConfig struct {
	Host string
	Port string
	Type string
}

func Setup(coreInitConfig *CoreInitConfig) *ApplicationConfig {
	appConfig := &ApplicationConfig{
		InitConfig: *coreInitConfig,
	}

	if coreInitConfig.Database {
		appConfig.setupDatabaseConfig()
	}

	appConfig.setupApplicationEnvironment()

	if coreInitConfig.Server {
		appConfig.setupServerConfig()
	}

	if coreInitConfig.Storage {
		appConfig.setupStorageConfig()
	}

	return appConfig
}

func (appConfig *ApplicationConfig) SetupModels(pool *pgxpool.Pool) {
	appConfig.Models = models.NewModels(pool)
}

func (appConfig *ApplicationConfig) setupApplicationEnvironment() {
	appConfig.Env = "ANONYMOUS_NATIVE_INSTANCE"
	appConfig.Service = "UNKNOWN_SERVICE"

	if os.Getenv("APPLICATION_ENV") != "" {
		appConfig.Env = os.Getenv("APPLICATION_ENV")
	}

	if os.Getenv("SERVICE_NAME") != "" {
		appConfig.Service = os.Getenv("SERVICE_NAME")
	}
}

func (appConfig *ApplicationConfig) setupDatabaseConfig() {
	appConfig.DatabaseConfig.Driver = "postgres"
	appConfig.DatabaseConfig.Host = WILDCARD_ADDR
	appConfig.DatabaseConfig.MaxIdleConns = 15
	appConfig.DatabaseConfig.MaxIdleTime = "15m"
	appConfig.DatabaseConfig.MaxOpenConns = 15
	appConfig.DatabaseConfig.Name = "antorus"
	appConfig.DatabaseConfig.Password = "pass1234"
	appConfig.DatabaseConfig.Sslmode = "disable"
	appConfig.DatabaseConfig.Port = "5432"
	appConfig.DatabaseConfig.User = "antorus"

	if os.Getenv("DB_DRIVER") != "" {
		appConfig.DatabaseConfig.Driver = os.Getenv("DB_DRIVER")
	}

	if os.Getenv("DB_HOST") != "" {
		appConfig.DatabaseConfig.Host = os.Getenv("DB_HOST")
	}

	if os.Getenv("DB_NAME") != "" {
		appConfig.DatabaseConfig.Name = os.Getenv("DB_NAME")
	}

	if os.Getenv("DB_PASSWORD") != "" {
		appConfig.DatabaseConfig.Password = os.Getenv("DB_PASSWORD")
	}

	if os.Getenv("DB_PORT") != "" {
		appConfig.DatabaseConfig.Port = os.Getenv("DB_PORT")
	}

	if os.Getenv("DB_SSLMODE") != "" {
		appConfig.DatabaseConfig.Sslmode = os.Getenv("DB_SSLMODE")
	}

	if os.Getenv("DB_USER") != "" {
		appConfig.DatabaseConfig.User = os.Getenv("DB_USER")
	}

	if os.Getenv("DB_MAX_IDLE_CONNS") != "" {
		maxIdleConns, err := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNS"))

		if err != nil {
			tmpLogger.Error(err.Error())
		}

		appConfig.DatabaseConfig.MaxIdleConns = maxIdleConns
	}

	if os.Getenv("DB_MAX_IDLE_TIME") != "" {
		appConfig.DatabaseConfig.MaxIdleTime = os.Getenv("DB_MAX_IDLE_TIME")
	}

	if os.Getenv("DB_MAX_OPEN_CONNS") != "" {
		maxOpenConns, err := strconv.ParseInt(os.Getenv("DB_MAX_OPEN_CONNS"), 10, 32)

		if err != nil {
			tmpLogger.Error(err.Error())

			maxOpenConns = 15
		}

		appConfig.DatabaseConfig.MaxOpenConns = int32(maxOpenConns)
	}
}

func (appConfig *ApplicationConfig) setupServerConfig() {
	appConfig.ServerConfig.Debug = "0"
	appConfig.ServerConfig.Host = WILDCARD_ADDR
	appConfig.ServerConfig.Port = "8080"
	appConfig.ServerConfig.TrustedOrigins = []string{"*"}

	if os.Getenv("DEBUG") != "" {
		appConfig.ServerConfig.Debug = os.Getenv("DEBUG")
	}

	if os.Getenv("HOST") != "" {
		appConfig.ServerConfig.Host = os.Getenv("HOST")
	}

	if os.Getenv("PORT") != "" {
		appConfig.ServerConfig.Port = os.Getenv("PORT")
	}

	if origins := os.Getenv("CORS_TRUSTED_ORIGINS"); origins != "" {
		appConfig.ServerConfig.TrustedOrigins = strings.Fields(origins)
	}
}

func (appConfig *ApplicationConfig) setupStorageConfig() {
	appConfig.StorageConfig.Host = WILDCARD_ADDR
	appConfig.StorageConfig.Port = "6379"
	appConfig.StorageConfig.Type = "REDIS"

	if os.Getenv("STORAGE_HOST") != "" {
		appConfig.StorageConfig.Host = os.Getenv("STORAGE_HOST")
	}

	if os.Getenv("STORAGE_PORT") != "" {
		appConfig.StorageConfig.Port = os.Getenv("STORAGE_PORT")
	}

	if os.Getenv("STORAGE_TYPE") != "" {
		appConfig.StorageConfig.Type = os.Getenv("STORAGE_TYPE")
	}
}
