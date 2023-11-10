package cmd

import (
	"os"
	"strconv"
)

var appConfig *AppConfig

type AppConfig struct {
	Api      *ApiConfig
	Postgres *PostgresConfig
}

type ApiConfig struct {
	ApiPort                    int
	ApiUrl                     string
	JWTSecret                  string
	JWTtokenExpiresDate        int
	JWTrefreshTokenExpiresDate int
}
type PostgresConfig struct {
	PgHost     string
	PgPort     string
	PgUser     string
	PgPassword string
	PgDbName   string
}

func getPGConfig() (*PostgresConfig, error) {
	return &PostgresConfig{
		PgHost:     os.Getenv("PG_HOST"),
		PgPort:     os.Getenv("PG_PORT"),
		PgUser:     os.Getenv("PG_USER"),
		PgPassword: os.Getenv("PG_PASSWORD"),
		PgDbName:   os.Getenv("PG_DB_NAME"),
	}, nil
}

func SetupAppConfig() error {

	ApiPort, err := strconv.Atoi(os.Getenv("API_PORT"))
	JWTtokenExpires, err := strconv.Atoi(os.Getenv("JWT_EXPIRES"))
	JWTrefreshTokenExpires, err := strconv.Atoi(os.Getenv("JWT_REFRESHTOKEN_EXPIRES"))
	apiConfig := &ApiConfig{ApiPort: ApiPort, ApiUrl: os.Getenv("API_URL"), JWTSecret: os.Getenv("JWT_SECRET"), JWTtokenExpiresDate: JWTtokenExpires, JWTrefreshTokenExpiresDate: JWTrefreshTokenExpires}
	pgConfig, err := getPGConfig()
	appConfig = &AppConfig{
		Api:      apiConfig,
		Postgres: pgConfig,
	}
	if err != nil {
		return err
	}

	return nil
}

func GetAppConfig() *AppConfig {
	return appConfig
}
