package main

import (
	"os"
	"sync"
	"time"

	"github.com/astrokiran/nimbus/internal/auth"
	"github.com/astrokiran/nimbus/internal/common/configs"
	"github.com/astrokiran/nimbus/internal/common/database"
	"github.com/astrokiran/nimbus/internal/common/log"
	"github.com/astrokiran/nimbus/internal/common/services"
	"github.com/astrokiran/nimbus/internal/consultant"
	users "github.com/astrokiran/nimbus/internal/user"
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
	logger     *zap.Logger
	config     config
	wg         sync.WaitGroup
	db         *database.Database
	auth       *auth.Auth
	consultant *consultant.Consultant
}

func run(logger *zap.Logger) error {
	var cfg config

	cfg.baseURL = configs.GetString("BASE_URL", "http://localhost:4444")
	cfg.httpPort = configs.GetInt("HTTP_PORT", 4444)
	cfg.db.dsn = configs.GetString("DB_DSN", "postgres:@localhost:5432/nimbus?sslmode=disable")
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

	// User
	users.InitUser(db, logger)
	userInstance, err := users.GetInstance()
	if err != nil {
		return err
	}

	// SMS Service
	smsService := services.NewSMSService("us-east-1")

	// Auth
	auth := auth.NewAuth(db, userInstance, smsService, logger)
	consultant := consultant.NewConsultant(db, auth, userInstance, smsService)

	app := &application{
		config:     cfg,
		db:         db,
		logger:     logger,
		auth:       auth,
		consultant: consultant,
	}

	return app.serveHTTP()
}

func main() {
	//Initialize Logger
	err := log.InitLogger(zap.Config{
		Level:         zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding:      "json",
		EncoderConfig: zap.NewProductionEncoderConfig(),
		OutputPaths: []string{
			"stdout",
		},
		ErrorOutputPaths: []string{
			"stderr",
		},
	})
	if err != nil {
		panic("failed to initialize logger: " + err.Error())
	}
	logger := log.GetLogger()
	defer logger.Sync()

	err = run(logger)
	if err != nil {
		trace := string(logger.Sync().Error())
		logger.Error(err.Error(), zap.Any("trace", trace))
		os.Exit(1)
	}
}
