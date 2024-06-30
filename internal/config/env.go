package config

import (
	"log"
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
	PGConfig               *PGConfig
}

func InitConfig() Config {

	pgconfig := PGConfig{
		DBHost:        getEnvStr("GOHTS_DB_HOST", "", true),
		DBPort:        getEnvStr("GOHTS_DB_PORT", "", true),
		DBUser:        getEnvStr("GOHTS_DB_USER", "", true),
		DBPassword:    getEnvStr("GOHTS_DB_PASS", "", true),
		DBName:        getEnvStr("GOHTS_DB_NAME", "", true),
		DBSchema:      getEnvStr("GOHTS_DB_SCHEMA", "", false),
		DBSSLMode:     getEnvStr("GOHTS_DB_SSLMODE", "require", true),
		DBConnTimeout: getEnvInt64("GOHTS_DB_CONNTIMEOUT", 0, true),
	}

	return Config{
		PublicHost:             getEnvStr("GOHTS_API_PUBLIC_HOST", "", false),
		ServerPort:             getEnvStr("GOHTS_API_PORT", "8080", true),
		JWTSecret:              getEnvStr("GOHTS_API_JWT_SECRET", "", true),
		JWTExpirationInSeconds: getEnvInt64("GOHTS_API_JWT_EXPIRATION_IN_SECONDS", 3600*24*7, false),
		PGConfig:               &pgconfig,
	}
}

func getEnvStr(key, fallback string, req bool) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	} else if !ok && req {
		log.Fatalf("Env key, %s, does not have a value", key)
	}
	return fallback
}

func getEnvInt64(key string, fallback int64, req bool) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			if req {
				log.Fatalf("Env key, %s, does not have a value", key)
			}
			return fallback
		}
		return i
	}
	if req {
		log.Fatalf("Env key, %s, does not have a value", key)
	}
	return fallback
}
