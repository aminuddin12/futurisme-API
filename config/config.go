package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App      AppConfig
	Database DatabaseConfig
	Redis    RedisConfig
	Security SecurityConfig
}

type AppConfig struct {
	Name  string `mapstructure:"APP_NAME"`
	Env   string `mapstructure:"APP_ENV"`
	Port  string `mapstructure:"APP_PORT"`
	Debug bool   `mapstructure:"APP_DEBUG"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"DB_HOST"`
	Port     string `mapstructure:"DB_PORT"`
	User     string `mapstructure:"DB_USER"`
	Pass     string `mapstructure:"DB_PASS"`
	Name     string `mapstructure:"DB_NAME"`
	SSLMode  string `mapstructure:"DB_SSLMODE"`
	TimeZone string `mapstructure:"DB_TIMEZONE"`
}

type RedisConfig struct {
	Host string `mapstructure:"REDIS_HOST"`
	Port string `mapstructure:"REDIS_PORT"`
	Pass string `mapstructure:"REDIS_PASS"`
	DB   int    `mapstructure:"REDIS_DB"`
}

type SecurityConfig struct {
	AppKey          string `mapstructure:"X_APP_KEY"`
	AppSecret       string `mapstructure:"X_APP_SECRET"`
	JWTSecret       string `mapstructure:"JWT_SECRET"`
	JWTExpiredHours int    `mapstructure:"JWT_EXPIRED_HOURS"` // Tipe data INT
}

// LoadConfig membaca file .env dan return struct Config
func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Warning: .env file not found, relying on system env vars")
	}

	var config Config
	if err := viper.Unmarshal(&config.App); err != nil {
		return nil, err
	}
	if err := viper.Unmarshal(&config.Database); err != nil {
		return nil, err
	}
	if err := viper.Unmarshal(&config.Redis); err != nil {
		return nil, err
	}
	if err := viper.Unmarshal(&config.Security); err != nil {
		return nil, err
	}

	return &config, nil
}
