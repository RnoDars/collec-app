package config

import (
	"os"
	"strconv"
)

// Config contient toute la configuration de l'application
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
	Kafka    KafkaConfig
}

// ServerConfig contient la configuration du serveur HTTP
type ServerConfig struct {
	Port string
	Env  string // development, staging, production
}

// DatabaseConfig contient la configuration de la base de données
type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// JWTConfig contient la configuration JWT
type JWTConfig struct {
	Secret           string
	AccessTokenTTL   int // en minutes
	RefreshTokenTTL  int // en heures
}

// KafkaConfig contient la configuration Kafka
type KafkaConfig struct {
	Brokers []string
	Enabled bool
}

// Load charge la configuration depuis les variables d'environnement
func Load() (*Config, error) {
	config := &Config{
		Server: ServerConfig{
			Port: getEnv("PORT", "8080"),
			Env:  getEnv("ENV", "development"),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnvAsInt("DB_PORT", 5432),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "postgres"),
			DBName:   getEnv("DB_NAME", "collec_app"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
		JWT: JWTConfig{
			Secret:          getEnv("JWT_SECRET", "change-me-in-production"),
			AccessTokenTTL:  getEnvAsInt("JWT_ACCESS_TTL", 15),
			RefreshTokenTTL: getEnvAsInt("JWT_REFRESH_TTL", 168), // 7 jours
		},
		Kafka: KafkaConfig{
			Brokers: []string{getEnv("KAFKA_BROKER", "localhost:9092")},
			Enabled: getEnvAsBool("KAFKA_ENABLED", false),
		},
	}

	return config, nil
}

// Helper functions pour lire les variables d'environnement

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := os.Getenv(key)
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}

func getEnvAsBool(key string, defaultValue bool) bool {
	valueStr := os.Getenv(key)
	if value, err := strconv.ParseBool(valueStr); err == nil {
		return value
	}
	return defaultValue
}
