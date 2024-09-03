package config

import (
	"fmt"
	"log"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresConfig struct {
	env Env
	dbs map[string]*gorm.DB
	mu  sync.RWMutex
}

func NewPostgresConfig(env Env) *PostgresConfig {
	return &PostgresConfig{
		env: env,
		dbs: make(map[string]*gorm.DB),
	}
}

func (p *PostgresConfig) isValid() bool {
	return p.env.DB_USER != "" && p.env.DB_PASS != "" && p.env.DB_HOST != "" && p.env.DB_PORT != ""
}

func (p *PostgresConfig) InitializeConnections(databaseNames []string) {
	for _, dbName := range databaseNames {
		dsn := p.BuildDBURL(dbName)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("Failed to connect to database %s: %v", dbName, err)
		}

		sqlDB, err := db.DB()
		if err != nil {
			log.Fatalf("Failed to configure database %s: %v", dbName, err)
		}

		sqlDB.SetMaxOpenConns(25)
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetConnMaxLifetime(30 * time.Minute)

		p.mu.Lock()
		p.dbs[dbName] = db
		p.mu.Unlock()
	}
}

func (p *PostgresConfig) BuildDBURL(databaseName string) string {
	if !p.isValid() {
		log.Fatalf("Missing one or more required environment variables: DB_USER, DB_PASS, DB_HOST, DB_PORT")
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		p.env.DB_USER, p.env.DB_PASS, p.env.DB_HOST, p.env.DB_PORT, databaseName)

	return dsn
}

func (p *PostgresConfig) GetDB(databaseName string) (*gorm.DB, error) {
	p.mu.RLock()
	db, exists := p.dbs[databaseName]
	p.mu.RUnlock()

	if !exists {
		return nil, fmt.Errorf("database connection for %s not initialized", databaseName)
	}

	return db, nil
}
