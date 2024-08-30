package config

import (
	"log"
	"path/filepath"

	"github.com/spf13/viper"
)

type Env struct {
	DATABASE_URL string `mapstructure:"DATABASE_URL"`
}

func NewEnv() *Env {

	projectRoot, err := filepath.Abs(filepath.Join(""))

	if err != nil {
		log.Fatalf("Error getting project root path: %v", err)
	}

	viper.SetConfigFile(filepath.Join(projectRoot, ".env"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	env := &Env{}
	viper.BindEnv("DATABASE_URL")

	if err := viper.Unmarshal(env); err != nil {
		log.Fatalf("Error unmarshalling config: %v", err)
	}

	return env
}
