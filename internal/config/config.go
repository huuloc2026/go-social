package config

import (
	"fmt"
	"log"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	App      AppConfig
	Database DatabaseConfig
	Redis    RedisConfig
	RabbitMQ RabbitMQConfig
}

type AppConfig struct {
	Name      string
	Env       string
	Port      string
	JWTSecret string
	JWTExpiry string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

type RabbitMQConfig struct {
	Host     string
	Port     string
	User     string
	Password string
}

var (
	once     sync.Once
	instance *Config
)

func Load() *Config {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./")
		viper.AddConfigPath("./internal/config/")
		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Error reading config file: %v", err)
		}

		instance = &Config{
			App: AppConfig{
				Name:      viper.GetString("app.name"),
				Env:       viper.GetString("app.env"),
				Port:      viper.GetString("app.port"),
				JWTSecret: viper.GetString("app.jwt_secret"),
				JWTExpiry: viper.GetString("app.jwt_expiry"),
			},
			Database: DatabaseConfig{
				Host:     viper.GetString("database.host"),
				Port:     viper.GetString("database.port"),
				User:     viper.GetString("database.user"),
				Password: viper.GetString("database.password"),
				Name:     viper.GetString("database.name"),
				SSLMode:  viper.GetString("database.sslmode"),
			},
			Redis: RedisConfig{
				Host:     viper.GetString("redis.host"),
				Port:     viper.GetString("redis.port"),
				Password: viper.GetString("redis.password"),
				DB:       viper.GetInt("redis.db"),
			},
			RabbitMQ: RabbitMQConfig{
				Host:     viper.GetString("rabbitmq.host"),
				Port:     viper.GetString("rabbitmq.port"),
				User:     viper.GetString("rabbitmq.user"),
				Password: viper.GetString("rabbitmq.password"),
			},
		}
	})

	return instance
}

func (c *DatabaseConfig) GetDSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.Name, c.SSLMode)
}
