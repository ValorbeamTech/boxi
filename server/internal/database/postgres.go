package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type DB struct {
    *sql.DB
}

// Config holds database configuration
type Config struct {
    Host     string
    Port     string
    User     string
    Password string
    DBName   string
    SSLMode  string
}

// NewPostgresDB creates a new PostgreSQL database connection
func NewPostgresDB(databaseURL string) (*sql.DB, error) {
    db, err := sql.Open("postgres", databaseURL)
    if err != nil {
        return nil, fmt.Errorf("failed to open database: %w", err)
    }

    // Configure connection pool
    db.SetMaxOpenConns(25)
    db.SetMaxIdleConns(5)
    db.SetConnMaxLifetime(5 * time.Minute)

    // Test the connection
    if err := db.Ping(); err != nil {
        return nil, fmt.Errorf("failed to ping database: %w", err)
    }

    return db, nil
}

// NewPostgresDBWithConfig creates a new PostgreSQL database connection using config
func NewPostgresDBWithConfig(cfg Config) (*sql.DB, error) {
    dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
        cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode)
    
    return NewPostgresDB(dsn)
}

// RunMigrations runs database migrations
func RunMigrations(db *sql.DB, migrationsPath string) error {
    driver, err := postgres.WithInstance(db, &postgres.Config{})
    if err != nil {
        return fmt.Errorf("failed to create migrate driver: %w", err)
    }

    m, err := migrate.NewWithDatabaseInstance(
        fmt.Sprintf("file://%s", migrationsPath),
        "postgres", 
        driver,
    )
    if err != nil {
        return fmt.Errorf("failed to create migrate instance: %w", err)
    }
    defer m.Close()

    if err := m.Up(); err != nil && err != migrate.ErrNoChange {
        return fmt.Errorf("failed to run migrations: %w", err)
    }

    return nil
}

// MigrateDown rolls back database migrations
func MigrateDown(db *sql.DB, migrationsPath string) error {
    driver, err := postgres.WithInstance(db, &postgres.Config{})
    if err != nil {
        return fmt.Errorf("failed to create migrate driver: %w", err)
    }

    m, err := migrate.NewWithDatabaseInstance(
        fmt.Sprintf("file://%s", migrationsPath),
        "postgres", 
        driver,
    )
    if err != nil {
        return fmt.Errorf("failed to create migrate instance: %w", err)
    }
    defer m.Close()

    if err := m.Down(); err != nil && err != migrate.ErrNoChange {
        return fmt.Errorf("failed to rollback migrations: %w", err)
    }

    return nil
}

// HealthCheck checks if the database is healthy
func HealthCheck(db *sql.DB) error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    
    return db.PingContext(ctx)
}

// Close closes the database connection
func (d *DB) Close() error {
    return d.DB.Close()
}