package app

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type HttpConfig struct {
	ReadTimeout  int
	WriteTimeout int
	IdleTimeout  int
}

type JWTConfig struct {
	Secret string
}

type Config struct {
	ConfigFile string
	Debug      bool
	Origins    []string
	DB         DatabaseConfig
	HTTP       HttpConfig
	JWT        JWTConfig
	Modules    []string
	Addr       string
	PprofAddr  string
}

var AppConfig Config

func LoadConfig() error {
	// Wczytywanie .env.default najpierw (może być zastąpione przez .env)
	_ = godotenv.Load(".env.default")

	// Wczytywanie .env (jeśli istnieje)
	_ = godotenv.Load()

	// Jeśli określono inny plik konfiguracyjny przez flagę
	if AppConfig.ConfigFile != "" {
		if err := godotenv.Overload(AppConfig.ConfigFile); err != nil {
			fmt.Printf("Warning: error loading specified config file %s: %v\n", AppConfig.ConfigFile, err)
		}
	}

	AppConfig.Addr = getEnv("ADDR", ":80")
	AppConfig.Debug = getEnvBool("DEBUG", false)
	AppConfig.PprofAddr = getEnv("PPROF_ADDR", ":6061")

	AppConfig.DB.Host = getEnv("DB_HOST", "localhost")
	AppConfig.DB.Port = getEnvInt("DB_PORT", 5432)
	AppConfig.DB.User = getEnv("DB_USER", "app")
	AppConfig.DB.Password = getEnv("DB_PASSWORD", "app")
	AppConfig.DB.DBName = getEnv("DB_NAME", "app")
	AppConfig.DB.SSLMode = getEnv("DB_SSLMODE", "disable")

	AppConfig.HTTP.ReadTimeout = getEnvInt("HTTP_READ_TIMEOUT", 5)
	AppConfig.HTTP.WriteTimeout = getEnvInt("HTTP_WRITE_TIMEOUT", 10)
	AppConfig.HTTP.IdleTimeout = getEnvInt("HTTP_IDLE_TIMEOUT", 10)

	AppConfig.JWT.Secret = getEnv("JWT_SECRET", "")

	modulesStr := getEnv("MODULES", "")
	if modulesStr != "" {
		AppConfig.Modules = strings.Split(modulesStr, ",")
	}

	originsStr := getEnv("ORIGINS", "")
	if originsStr != "" {
		AppConfig.Origins = strings.Split(originsStr, ",")
	}

	return nil
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

func getEnvInt(key string, defaultValue int) int {
	valueStr, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultValue
	}
	return value
}

func getEnvBool(key string, defaultValue bool) bool {
	valueStr, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	value, err := strconv.ParseBool(valueStr)
	if err != nil {
		return defaultValue
	}
	return value
}
