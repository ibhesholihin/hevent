package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config ...
type Config struct {
	DatabaseURL         string
	CacheURL            string
	LoggerLevel         string
	ContextTimeout      int
	JWTSecretKey        string
	DB_PORT             string
	DB_HOST             string
	DB_NAME             string
	DB_USER             string
	DB_PASS             string
	MIDTRANS_SERVER_KEY string
	MIDTRANS_CLIENT_KEY string
	BaseURL             string
}

// LoadConfig will load config from environment variable
func LoadConfig() (config *Config) {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	databaseURL := os.Getenv("DATABASE_URL")
	cacheURL := os.Getenv("CACHE_URL")
	loggerLevel := os.Getenv("LOGGER_LEVEL")
	contextTimeout, _ := strconv.Atoi(os.Getenv("CONTEXT_TIMEOUT"))
	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")

	midtrans_SERVER_KEY := os.Getenv("MIDTRANS_SERVER_KEY")
	midtrans_CLIENT_KEY := os.Getenv("MIDTRANS_CLIENT_KEY=SB-Mid-client")

	db_PORT := os.Getenv("DB_PORT")
	db_HOST := os.Getenv("DB_HOST")
	db_NAME := os.Getenv("DB_NAME")
	db_USER := os.Getenv("DB_USER")
	db_PASS := os.Getenv("DB_PASS")

	base_url := os.Getenv("BASE_URI")

	return &Config{
		DatabaseURL:    databaseURL,
		CacheURL:       cacheURL,
		LoggerLevel:    loggerLevel,
		ContextTimeout: contextTimeout,
		JWTSecretKey:   jwtSecretKey,

		DB_PORT: db_PORT,
		DB_HOST: db_HOST,
		DB_NAME: db_NAME,
		DB_USER: db_USER,
		DB_PASS: db_PASS,

		MIDTRANS_SERVER_KEY: midtrans_SERVER_KEY,
		MIDTRANS_CLIENT_KEY: midtrans_CLIENT_KEY,

		BaseURL: base_url,
	}
}
