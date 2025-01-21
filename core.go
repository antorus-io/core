package core

import (
	"log/slog"
	"os"

	"github.com/antorus-io/core/config"
	"github.com/antorus-io/core/database"
	"github.com/antorus-io/core/logs"
	"github.com/antorus-io/core/storage"
)

func Init() {
	app := config.Setup()

	initDatabase(app)
	initLogger(app)
	initStorage(app)

	logs.Logger.Info("Successfully initialized Core module")
}

func initDatabase(app *config.ApplicationConfig) {
	err := database.CreateDatabase(app.DatabaseConfig)
	tmpLogger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	if err != nil {
		tmpLogger.Error("Error initializing database", "error", err)

		os.Exit(1)
	}

	if err != nil {
		tmpLogger.Error("Error opening database", "error", err.Error())

		os.Exit(1)
	}

	// TODO: move this to individual services?
	app.SetupModels(database.DatabaseInstance.GetPool())

	tmpLogger.Info("Database connection pool established")
}

func initLogger(app *config.ApplicationConfig) {
	logs.CreateLogger(app)
}

func initStorage(app *config.ApplicationConfig) {
	err := storage.CreateStorage(app.StorageConfig)

	if err != nil {
		logs.Logger.Error("Error initializing storage", "error", err)

		os.Exit(1)
	}

	logs.Logger.Info("Storage connection successfully initialized")
}
