package config

import (
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

func (p *PostgresConfig) Client() *gorm.DB {
	if p.db != nil {
		return p.db
	}

	dsn := p.env.DATABASE_URL
	var err error
	p.db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	return p.db
}

func (p *PostgresConfig) Close() {
	sqlDB, err := p.Client().DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}

	err = sqlDB.Close()
	if err != nil {
		log.Fatalf("Failed to close database: %v", err)
	}
}

func (p *PostgresConfig) Migrate(models ...interface{}) {
	db := p.Client()
	err := db.AutoMigrate(models...)
	if err != nil {
		log.Fatalf("Failed to migrate database schema: %v", err)
	}
}
