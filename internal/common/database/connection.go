package database

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Database struct {
	Conn *sqlx.DB
}

type Config struct {
	DSN             string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

func NewDatabase(cfg Config) (*Database, error) {
	db, err := sqlx.Open("postgres", cfg.DSN)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Set the maximum number of open connections to the database.
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	// Set the maximum number of idle connections in the pool.
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	// Set the maximum lifetime of a connection.
	db.SetConnMaxLifetime(cfg.ConnMaxLifetime)

	return &Database{Conn: db}, nil
}

// Close gracefully shuts down the database connection.
func (d *Database) Close() error {
	if d.Conn != nil {
		return d.Conn.Close()
	}
	return nil
}
