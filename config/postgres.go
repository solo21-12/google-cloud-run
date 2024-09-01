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

func (p *PostgresConfig) isValid() bool {
	return p.env.DB_USER != "" && p.env.DB_PASS != "" && p.env.DB_HOST != "" && p.env.DB_PORT != ""
}

func (p *PostgresConfig) BuildDBURL(databaseName string) string {
	if !p.isValid() {
		log.Fatalf("Missing one or more required environment variables: DB_USER, DB_PASS, DB_HOST, DB_PORT")
	}

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

func (p *PostgresConfig) Close(databaseName string) {
	sqlDB, err := p.Client(databaseName).DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}

	err = sqlDB.Close()
	if err != nil {
		log.Fatalf("Failed to close database: %v", err)
	}
}

func (p *PostgresConfig) Migrate(databaseName string, models ...interface{}) {
	db := p.Client(databaseName)
	err := db.AutoMigrate(models...)
	if err != nil {
		log.Fatalf("Failed to migrate database schema: %v", err)
	}
}
