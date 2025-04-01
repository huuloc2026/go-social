package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBHost            string        `mapstructure:"DB_HOST"`
	DBPort            string        `mapstructure:"DB_PORT"`
	DBUser            string        `mapstructure:"DB_USER"`
	DBPassword        string        `mapstructure:"DB_PASSWORD"`
	DBName            string        `mapstructure:"DB_NAME"`
	Port              string        `mapstructure:"PORT"`
	JWTSecret         string        `mapstructure:"JWT_SECRET"`
	JWTExpiration     time.Duration `mapstructure:"JWT_EXPIRATION"`
	RefreshExpiration time.Duration `mapstructure:"REFRESH_EXPIRATION"`
	RedisHost         string        `mapstructure:"RedisHost"`
	RedisPort         string        `mapstructure:"RedisPort"`
	RedisPassword     string        `mapstructure:"RedisPassword"`
	RedisDB           string        `mapstructure:"RedisDB"`
	RedisURL          string        `mapstructure:"REDIS_URL"`
	SMTPHost          string        `mapstructure:"SMTP_HOST"`
	SMTPPort          int           `mapstructure:"SMTP_PORT"`
	SMTPUser          string        `mapstructure:"SMTP_USER"`
	SMTPPassword      string        `mapstructure:"SMTP_PASSWORD"`
	EmailFrom         string        `mapstructure:"EMAIL_FROM"`
	RateLimit         int           `mapstructure:"RATE_LIMIT"`
	RateLimitWindow   time.Duration `mapstructure:"RATE_LIMIT_WINDOW"`
	Environment       string        `mapstructure:"ENVIRONMENT"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	// Set defaults
	viper.SetDefault("PORT", "8080")
	viper.SetDefault("JWT_EXPIRATION", "15m")
	viper.SetDefault("REFRESH_EXPIRATION", "168h") // 7 days
	viper.SetDefault("RATE_LIMIT", 100)
	viper.SetDefault("RATE_LIMIT_WINDOW", "1m")
	viper.SetDefault("ENVIRONMENT", "development")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
