package core

import (
	"context"
	"log/slog"
	"os"

	"github.com/antorus-io/core/config"
	"github.com/antorus-io/core/database"
	"github.com/antorus-io/core/events"
	"github.com/antorus-io/core/logs"
	"github.com/antorus-io/core/server"
	"github.com/antorus-io/core/storage"
)

func Init(coreInitConfig *config.CoreInitConfig) *config.ApplicationConfig {
	appConfig := config.Setup(coreInitConfig)

	if coreInitConfig.Database {
		initDatabase(appConfig)
	}

	if coreInitConfig.Logger {
		initLogger(appConfig)
	}

	if coreInitConfig.Storage {
		initStorage(appConfig)
	}

	if coreInitConfig.Database && coreInitConfig.Storage {
		initEventRegistry(appConfig)
	}

	return appConfig
}

func StartServer(appConfig *config.ApplicationConfig) error {
	server.NewServer(appConfig)

	if err := server.ServerInstance.Serve(); err != nil {
		logs.Logger.Error("Server error", "error", err)

		return err
	}

	return nil
}

func initDatabase(appConfig *config.ApplicationConfig) {
	err := database.CreateDatabase(appConfig.DatabaseConfig)
	tmpLogger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	if err != nil {
		tmpLogger.Error("Error initializing database", "error", err)

		os.Exit(1)
	}

	if err != nil {
		tmpLogger.Error("Error opening database", "error", err.Error())

		os.Exit(1)
	}

	appConfig.SetupModels(database.DatabaseInstance.GetPool())

	tmpLogger.Info("Database connection successfully initialized")
}

func initEventRegistry(appConfig *config.ApplicationConfig) {
	err := events.InitEventRegistry(context.Background(), database.DatabaseInstance.GetPool())

	if err != nil {
		logs.Logger.Error("Error initializing EventRegistry", "error", err)
	}

	appConfig.Events = events.GetEventRegistry().AllEvents()
}

func initLogger(appConfig *config.ApplicationConfig) {
	logs.CreateLogger(appConfig)
}

func initStorage(appConfig *config.ApplicationConfig) {
	err := storage.CreateStorage(appConfig.StorageConfig)

	if err != nil {
		logs.Logger.Error("Error initializing storage", "error", err)

		os.Exit(1)
	}

	logs.Logger.Info("Storage connection successfully initialized")
}
