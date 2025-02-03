package database

import (
	"context"
	"fmt"
	"time"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Dialect struct {
	Int    func(value int64) postgres.IntegerExpression
	String func(value string) postgres.StringExpression
	UUID   func(value fmt.Stringer) postgres.StringExpression
}

var PostgresDialect = Dialect{
	Int:    postgres.Int,
	String: postgres.String,
	UUID:   postgres.UUID,
}

type Database struct {
	Conn    *sqlx.DB
	Dialect Dialect
}

type Config struct {
	DSN             string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

func NewDatabase(cfg Config) (*Database, error) {
	fmt.Println("Connecting to database...", cfg.DSN)
	ctx, cancel := context.WithTimeout(context.Background(), cfg.ConnMaxLifetime)
	defer cancel()

	db, err := sqlx.ConnectContext(ctx, "postgres", "postgres://"+cfg.DSN)
	if err != nil {
		return nil, err
	}

	// Set the maximum number of open connections to the database.
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	// Set the maximum number of idle connections in the pool.
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	// Set the maximum lifetime of a connection.
	db.SetConnMaxLifetime(cfg.ConnMaxLifetime)

	return &Database{Conn: db, Dialect: PostgresDialect}, nil
}

// Close gracefully shuts down the database connection.
func (d *Database) Close() error {
	if d.Conn != nil {
		return d.Conn.Close()
	}
	return nil
}
