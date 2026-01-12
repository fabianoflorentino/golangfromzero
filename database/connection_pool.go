package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// ConnectionPoolConfig represents a database pool configuration
type ConnectionPoolConfig struct {
	Host                  string
	Port                  string
	User                  string
	Password              string
	Database              string
	SSLmode               string
	MaxConnections        int32
	MinConnections        int32
	MaxConnectionLifetime time.Duration
	MaxConnectionIdletime time.Duration
	HealthcheckPeriod     time.Duration
}

// LoadConnectionPoolConfig load the configurations from environment variables.
func LoadConnectionPoolConfig() ConnectionPoolConfig {
	return ConnectionPoolConfig{
		Host:                  getEnv("DB_HOST", "localhost"),
		Port:                  getEnv("DB_PORT", "5432"),
		User:                  getEnv("DB_USER", "postgres"),
		Password:              getEnv("DB_PASSWORD", "postgres"),
		Database:              getEnv("DB_NAME", "exampledb"),
		SSLmode:               getEnv("DB_SSLMODE", "disable"),
		MaxConnections:        getEnvAsInt32("DB_MAX_CONNECTIONS", 25),
		MinConnections:        getEnvAsInt32("DB_MIN_CONNEXTIONS", 5),
		MaxConnectionLifetime: getEnvAsDuration("DB_MAX_CONNECTION_LIFETIME", 5*time.Minute),
		MaxConnectionIdletime: getEnvAsDuration("DB_MAX_CONNECTION_IDLETIME", 1*time.Minute),
		HealthcheckPeriod:     getEnvAsDuration("DB_HEALTHCHECK_PERIOD", 1*time.Minute),
	}
}

// NewConnectionPool creates a new database connection pool
func NewConnectionPool(ctx context.Context, cfg ConnectionPoolConfig) (*pgxpool.Pool, error) {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.SSLmode)

	connectionPoolConfig, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		return nil, fmt.Errorf("unable to parse database config: %w", err)
	}

	connectionPoolConfig.MaxConns = cfg.MaxConnections
	connectionPoolConfig.MinConns = cfg.MinConnections
	connectionPoolConfig.MaxConnLifetime = cfg.MaxConnectionLifetime
	connectionPoolConfig.MaxConnIdleTime = cfg.MaxConnectionIdletime

	connectionPool, err := pgxpool.NewWithConfig(ctx, connectionPoolConfig)
	if err != nil {
		return nil, fmt.Errorf("unable to create database pool: %w", err)
	}

	if err := connectionPool.Ping(ctx); err != nil {
		connectionPool.Close()
		return nil, fmt.Errorf("unable to ping database: %w", err)
	}

	log.Printf("database connection pool created successfully (max: %d, min: %d)", cfg.MaxConnections, cfg.MinConnections)

	return connectionPool, nil
}

// getEnv get value from environment variable or set a default value
func getEnv(env, defaultValue string) string {
	if value := os.Getenv(env); value != "" {
		return value
	}

	return defaultValue
}

// getEnvAsInt32 get value from environment variables converting string to int32 or set a default value
func getEnvAsInt32(env string, defaultValue int32) int32 {
	if value := os.Getenv(env); value != "" {
		var intValue int32

		fmt.Sscanf(value, "%d", &intValue)
		return intValue
	}

	return defaultValue
}

// getEnvAsDuration get value from environment variables converting string to time.Duration or set a default value
func getEnvAsDuration(env string, defaultValue time.Duration) time.Duration {
	if value := os.Getenv(env); value != "" {
		duration, err := time.ParseDuration(value)
		if err == nil {
			return duration
		}
	}

	return defaultValue
}
