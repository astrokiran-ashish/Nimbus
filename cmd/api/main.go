package main

import (
	"os"
	"sync"
	"time"

	"github.com/astrokiran/nimbus/internal/common/configs"
	"github.com/astrokiran/nimbus/internal/common/database"
	"go.uber.org/zap"
)

type config struct {
	baseURL  string
	httpPort int
	db       struct {
		dsn         string
		automigrate bool
	}
	idleTimeout    time.Duration
	readTimeout    time.Duration
	writeTimeout   time.Duration
	shutdownPeriod time.Duration
}

type application struct {
	logger *zap.Logger
	config config
	wg     sync.WaitGroup
	db     *database.Database
}

func run(logger *zap.Logger) error {
	var cfg config

	cfg.baseURL = configs.GetString("BASE_URL", "http://localhost:4444")
	cfg.httpPort = configs.GetInt("HTTP_PORT", 4444)
	cfg.db.dsn = configs.GetString("DB_DSN", "user:pass@localhost:5432/db")
	cfg.db.automigrate = configs.GetBool("DB_AUTOMIGRATE", true)
	databaseConfig := database.Config{
		DSN:             cfg.db.dsn,
		MaxOpenConns:    configs.GetInt("DB_MAX_OPEN_CONNS", 10),
		MaxIdleConns:    configs.GetInt("DB_MAX_IDLE_CONNS", 10),
		ConnMaxLifetime: 10 * time.Second,
	}
	db, err := database.NewDatabase(databaseConfig)
	if err != nil {
		return err
	}
	defer db.Close()

	app := &application{
		config: cfg,
		db:     db,
		logger: logger,
	}

	return app.serveHTTP()
}

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	err := run(logger)
	if err != nil {
		trace := string(logger.Sync().Error())
		logger.Error(err.Error(), zap.Any("trace", trace))
		os.Exit(1)
	}
}
