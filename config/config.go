package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Logger 		LoggerConfig 	`mapstructure:",squash"`
	HTTP 		HTTPConfig 		`mapstructure:",squash"`
	Database 	DatabaseConfig	`mapstructure:",squash"`
}

type LoggerConfig struct {
	Level 	string `mapstructure:"LOGGER_LEVEL"`
	File 	string `mapstructure:"LOGGER_FILE"`
}

type HTTPConfig struct {
	Host 	string 	`mapstructure:"HOST"`
	Port 	int64 	`mapstructure:"PORT"`
}

type DatabaseConfig struct {
	Postgres 	PostgresConfig 	`mapstructure:",squash"`
	Driver 		string 			`mapstructure:"DB_DRIVER"`
}

type PostgresConfig struct {
	Host 		string 	`mapstructure:"DB_HOST"`
	Port 		int64 	`mapstructure:"DB_PORT"`
	User 		string 	`mapstructure:"DB_USER"`
	Password 	string 	`mapstructure:"DB_PASSWORD"`
	Database 	string 	`mapstructure:"DB_NAME"`
}

func LoadConfig(path string) (*Config, error) {
	var config Config

	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err:= viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}