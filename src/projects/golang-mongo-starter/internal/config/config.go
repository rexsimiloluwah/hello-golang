package config

import (
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

// For reading the environment variables
type Settings struct {
	Env          string `mapstructure:"ENV"`
	DbName       string `mapstructure:"DB_NAME"`
	DbHost       string `mapstructure:"DB_HOST"`
	DbPort       string `mapstructure:"DB_PORT"`
	DbUser       string `mapstructure:"DB_USER"`
	DbPass       string `mapstructure:"DB_PASS"`
	DbProdUri    string `mapstructure:"DB_PROD_URI"`
	JwtSecretKey string `mapstructure:"JWT_SECRET_KEY"`
	JwtExpiresIn string `mapstructure:"JWT_EXPIRES_IN"`
}

func New() *Settings {
	var cfg Settings

	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		log.Errorf("Error reading environment variables: %v", err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatalf("Error unmarshaling to config struct: %v", err)
	}

	return &cfg
}
