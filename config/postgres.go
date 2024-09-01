package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresConfig struct {
	env Env
	db  *gorm.DB
}

func NewPostgresConfig(env Env) *PostgresConfig {
	return &PostgresConfig{env: env}
}

func (p *PostgresConfig) BuildDBURL(databaseName string) string {
	// Debugging: Log environment variables
	log.Printf("DB_USER: %s, DB_PASS: %s, DB_HOST: %s, DB_PORT: %s, DB_NAME: %s",
		p.env.DB_USER, p.env.DB_PASS, p.env.DB_HOST, p.env.DB_PORT, databaseName)

	// Check if any of the environment variables are missing
	if p.env.DB_USER == "" || p.env.DB_PASS == "" || p.env.DB_HOST == "" || p.env.DB_PORT == "" {
		log.Fatalf("Missing one or more required environment variables: DB_USER, DB_PASS, DB_HOST, DB_PORT")
	}

	// Construct the PostgreSQL DSN (Data Source Name)
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		p.env.DB_USER, p.env.DB_PASS, p.env.DB_HOST, p.env.DB_PORT, databaseName)

	return dsn
}

func (p *PostgresConfig) Client(databaseName string) *gorm.DB {
	if p.db != nil {
		return p.db
	}

	dsn := p.BuildDBURL(databaseName)
	var err error
	p.db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	return p.db
}

func (p *PostgresConfig) Close() {
	sqlDB, err := p.Client(p.env.DB_NAME).DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}

	err = sqlDB.Close()
	if err != nil {
		log.Fatalf("Failed to close database: %v", err)
	}
}

func (p *PostgresConfig) Migrate(models ...interface{}) {
	db := p.Client(p.env.DB_NAME)
	err := db.AutoMigrate(models...)
	if err != nil {
		log.Fatalf("Failed to migrate database schema: %v", err)
	}
}
