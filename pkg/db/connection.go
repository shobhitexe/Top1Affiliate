package db

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Connection struct{}

func NewConnection() *Connection {
	return &Connection{}
}

func (c *Connection) ConnectToPostgres(dsn string, maxConns, minConns int, maxIdleTime string) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Printf("Failed to parse database connection string: %v", err)
		return nil, err
	}

	config.MaxConns = int32(maxConns)
	config.MinConns = int32(minConns)

	if maxIdleTime != "" {
		idleTime, err := time.ParseDuration(maxIdleTime)
		if err != nil {
			log.Printf("Invalid maxIdleTime value: %v", err)
			return nil, err
		}
		config.MaxConnIdleTime = idleTime
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Printf("Failed to create PostgreSQL connection pool: %v", err)
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		log.Printf("Failed to ping PostgreSQL database: %v", err)
		pool.Close()
		return nil, err
	}

	go c.logPoolStats(pool)

	log.Println("Successfully connected to PostgreSQL")
	return pool, nil
}

func (c *Connection) logPoolStats(pool *pgxpool.Pool) {
	ticker := time.NewTicker(60 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		stats := pool.Stat()
		log.Printf("[Postgres Pool Stats] TotalConns: %d, IdleConns: %d, UsedConns: %d, MaxConns: %d",
			stats.TotalConns(), stats.IdleConns(), stats.AcquiredConns(), stats.MaxConns())
	}
}
