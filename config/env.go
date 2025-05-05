package config

import (
	"github.com/joho/godotenv"
	"log"
	"log/slog"
	"os"
	"strconv"
)

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Println("failed load env file")
	}
	log.Println("env file is load")
}

type Config struct {
	DbUrl string
}

func NewConfig() *Config {
	return &Config{
		DbUrl: getString("DATABASE_URL", ""),
	}
}

type LogConfig struct {
	Level slog.Level
}

func NewLogConfig() *LogConfig {
	return &LogConfig{
		Level: slog.Level(getInt("LEVEL_LOG", -4)),
	}
}

func getString(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	i, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return i
}

func getBool(key string, defaultValue bool) bool {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	b, err := strconv.ParseBool(value)
	if err != nil {
		return defaultValue
	}
	return b
}
