package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type DBConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
}

type Config struct {
	PostgresDB DBConfig `mapstructure:"database"`
}

func LoadConfig() *Config {
	// Initialize Viper
	viper.SetConfigName("config") // Name of config file (without extension)
	viper.SetConfigType("yaml")   // Config file type
	viper.AddConfigPath("./")     // Look for config in the working directory
	viper.AddConfigPath("./internal/config/")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("No config file found, using defaults and environment variables")
		} else {
			panic(fmt.Errorf("fatal error reading config file: %v", err))
		}
	}
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Unable to decode config: %v", err)
	}
	return &cfg
}
