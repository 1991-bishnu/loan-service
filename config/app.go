package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type AppConfig struct {
	Server struct {
		Address string `mapstructure:"address"`
		Mode    string `mapstructure:"mode"`
		Timeout int    `mapstructure:"timeout"`
	} `mapstructure:"server"`

	Database struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Name     string `mapstructure:"name"`
	} `mapstructure:"database"`

	Log struct {
		Level string `mapstructure:"level"`
	} `mapstructure:"log"`

	Auth struct {
		Token string `mapstructure:"token"`
	} `mapstructure:"auth"`
}

func LoadConfig() (*AppConfig, error) {
	viper.SetConfigName("dev")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")

	// Read from environment variables if available
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var config AppConfig
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}
	return &config, nil
}
