package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/astrokiran/nimbus/internal/auth"
	"github.com/astrokiran/nimbus/internal/common/configs"
	"github.com/astrokiran/nimbus/internal/common/database"
	"github.com/astrokiran/nimbus/internal/common/log"
	orchestrationengine "github.com/astrokiran/nimbus/internal/common/orchestration-engine"
	"github.com/astrokiran/nimbus/internal/common/services"
	"github.com/astrokiran/nimbus/internal/consultant"
	"github.com/astrokiran/nimbus/internal/consultation"
	"github.com/astrokiran/nimbus/internal/notification"
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
	jwtSecret      string
	jwtExpiry      time.Duration
	refreshExpiry  time.Duration
	taskQueue      string
	temporalHost   string
	temporalPort   int
}

type application struct {
	logger       *zap.Logger
	config       config
	wg           sync.WaitGroup
	db           *database.Database
	auth         *auth.Auth
	consultant   *consultant.Consultant
	consultation *consultation.Consultation
}

func run(logger *zap.Logger) error {
	var cfg config

	cfg.baseURL = configs.GetString("BASE_URL", "http://localhost:4444")
	cfg.httpPort = configs.GetInt("HTTP_PORT", 4444)
	cfg.db.dsn = configs.GetString("DB_DSN", "postgres:@localhost:5432/nimbus?sslmode=disable")
	cfg.db.automigrate = configs.GetBool("DB_AUTOMIGRATE", true)
	cfg.jwtSecret = configs.GetString("JWT_SECRET", "secret")
	cfg.jwtExpiry = time.Duration(configs.GetInt("JWT_EXPIRY_MINS", 15))
	cfg.refreshExpiry = time.Duration(configs.GetInt("JWT_REFRESH_TOKEN_EXPIRY_DAYS", 30))
	cfg.taskQueue = configs.GetString("TASK_QUEUE", "nimbus")
	cfg.temporalHost = configs.GetString("TEMPORAL_HOST", "localhost")
	cfg.temporalPort = configs.GetInt("TEMPORAL_PORT", 7233)

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

	temporalHostPort := fmt.Sprintf("%s:%s", cfg.temporalHost, strconv.Itoa(cfg.temporalPort))
	fmt.Printf("Temporal Host: %s\n", temporalHostPort)
	temporalClient, err := CreateTemporalClient(temporalHostPort)
	if err != nil {
		return err
	}

	engine := orchestrationengine.NewOrchestrationEngine(temporalClient)

	// User
	users.InitUser(db, logger)
	userInstance, err := users.GetInstance()
	if err != nil {
		return err
	}

	// SMS Service
	smsService := services.NewSMSService("us-east-1")

	// Auth
	auth := auth.NewAuth(db, userInstance, smsService, logger, cfg.jwtSecret, cfg.jwtExpiry, cfg.refreshExpiry)
	consultant := consultant.NewConsultant(db, auth, userInstance, smsService)
	consultation := consultation.NewConsultation(db, engine, cfg.taskQueue)
	notification := notification.NewNotification(db, logger)

	app := &application{
		config:       cfg,
		db:           db,
		logger:       logger,
		auth:         auth,
		consultant:   consultant,
		consultation: consultation,
	}

	// Start Temporal worker in a new goroutine
	if err := app.startWorker(temporalClient, cfg.taskQueue, notification); err != nil {
		logger.Error("failed to start Temporal worker", zap.Error(err))
		os.Exit(1)
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
