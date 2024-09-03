package config

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	DB_USER    string   `mapstructure:"DB_USER"`
	DB_PASS    string   `mapstructure:"DB_PASS"`
	DB_HOST    string   `mapstructure:"DB_HOST"`
	DB_PORT    string   `mapstructure:"DB_PORT"`
	JWT_SECRET string   `mapstructure:"JWT_SECRET"`
	DB_NAMES   string `mapstructure:"DB_NAMES"`
}

func NewEnv() *Env {
	viper.AutomaticEnv()
	env := &Env{}

	viper.BindEnv("DB_USER")
	viper.BindEnv("DB_PASS")
	viper.BindEnv("DB_HOST")
	viper.BindEnv("DB_PORT")
	viper.BindEnv("JWT_SECRET")
	viper.BindEnv("DB_NAMES")

	if err := viper.Unmarshal(env); err != nil {
		log.Fatalf("Error unmarshalling config: %v", err)
	}

	return env
}
