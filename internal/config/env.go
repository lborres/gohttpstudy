package config

import (
	"os"
	"strconv"
)

type PGConfig struct {
	DBUser        string
	DBPassword    string
	DBHost        string
	DBPort        string
	DBName        string
	DBSchema      string
	DBSSLMode     string
	DBConnTimeout int64
}

type Config struct {
	PublicHost             string
	ServerPort             string
	JWTSecret              string
	JWTExpirationInSeconds int64

	// TODO Should this be a pointer? How do I check memory usage?
	PGConfig *PGConfig
}

func InitConfig() Config {

	pgconfig := PGConfig{
		DBHost:        getEnvStr("LOCAL_DB_HOST", "127.0.0.1"),
		DBPort:        getEnvStr("LOCAL_DB_PORT", "5432"),
		DBUser:        getEnvStr("LOCAL_DB_USER", ""),
		DBPassword:    getEnvStr("LOCAL_DB_PASS", ""),
		DBName:        getEnvStr("LOCAL_DB_NAME", "localdev"),
		DBSchema:      getEnvStr("LOCAL_DB_SCHEMA", ""),
		DBSSLMode:     getEnvStr("LOCAL_DB_SSLMODE", "require"),
		DBConnTimeout: getEnvInt64("LOCAL_DB_CONNTIMEOUT", 0),
	}

	return Config{
		PublicHost:             getEnvStr("PUBLIC_HOST", ""),
		ServerPort:             getEnvStr("SERVER_PORT", "8080"),
		JWTSecret:              getEnvStr("JWT_SECRET", ""),
		JWTExpirationInSeconds: getEnvInt64("JWT_EXPIRATION_IN_SECONDS", 3600*24*7),
		PGConfig:               &pgconfig,
	}
}

func getEnvStr(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvInt64(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}
		return i
	}
	return fallback
}
