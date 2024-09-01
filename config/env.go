package config

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	DB_USER    string `mapstructure:"DB_USER"`
	DB_PASS    string `mapstructure:"DB_PASS"`
	DB_HOST    string `mapstructure:"DB_HOST"`
	DB_PORT    string `mapstructure:"DB_PORT"`
	DB_NAME    string `mapstructure:"DB_NAME"`
	JWT_SECRET string `mapstructure:"JWT_SECRET"`
}

func NewEnv() *Env {
	// projectRoot, err := filepath.Abs(filepath.Join(""))
	// if err != nil {
	// 	log.Fatalf("Error getting project root path: %v", err)
	// }

	// viper.SetConfigFile(filepath.Join(projectRoot, ".env"))
	viper.AutomaticEnv()

	// if err := viper.ReadInConfig(); err != nil {
	// 	log.Fatalf("Error reading config file: %v", err)
	// }

	env := &Env{}
	viper.BindEnv("DB_USER")
	viper.BindEnv("DB_PASS")
	viper.BindEnv("DB_HOST")
	viper.BindEnv("DB_PORT")
	viper.BindEnv("DB_NAME")
	viper.BindEnv("JWT_SECRET")

	if err := viper.Unmarshal(env); err != nil {
		log.Fatalf("Error unmarshalling config: %v", err)
	}

	return env
}
