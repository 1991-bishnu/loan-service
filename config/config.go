package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	ServerAddress    string `mapstructure:"server_address"`
	ServerMode       string `mapstructure:"server_mode"`
	ServerTimeout    int    `mapstructure:"server_timeout"`
	DatabaseHost     string `mapstructure:"database_host"`
	DatabasePort     int    `mapstructure:"database_port"`
	DatabaseUser     string `mapstructure:"database_user"`
	DatabasePassword string `mapstructure:"database_password"` // Insecure! Use secrets management.
	DatabaseName     string `mapstructure:"database_name"`
	LogLevel         string `mapstructure:"log_level"`
}

func LoadConfig() (*Config, error) {
	viper.AutomaticEnv()            // Read environment variables
	viper.SetConfigName("config")   // Name of the config file (without extension)
	viper.SetConfigType("env")      // Type of the config file (yaml, json, toml, etc.)
	viper.AddConfigPath(".")        // Path to the config file (e.g., current directory)
	viper.AddConfigPath("./config") // Check config directory as well

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}
	return &config, nil
}
