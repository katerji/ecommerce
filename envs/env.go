package envs

import (
	"github.com/joho/godotenv"
	"os"
	"path/filepath"
)

const (
	defaultDBHost           = "127.0.0.1"
	defaultDBPort           = "3306"
	defaultDBUser           = "root"
	defaultDBPassword       = "root"
	defaultDBDatabase       = "app"
	defaultJWTSecret        = "dummy"
	defaultJWTExpiry        = "86400"
	defaultRefreshJWTSecret = "dummy_refresh"
	defaultRefreshJWTExpiry = "1209600"
)

type env struct {
	jwtSecret        string
	jwtExpiry        string
	jwtRefreshSecret string
	jwtRefreshExpiry string
	dbHost           string
	dbPassword       string
	dbUser           string
	dbPort           string
	dbName           string
}

func getFromEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value != "" {
		return value
	}
	return defaultValue
}

func newEnv() *env {
	loadEnv()
	return &env{
		jwtSecret:        getFromEnv("JWT_SECRET", defaultJWTSecret),
		jwtExpiry:        getFromEnv("JWT_EXPIRY", defaultJWTExpiry),
		jwtRefreshSecret: getFromEnv("JWT_REFRESH_SECRET", defaultRefreshJWTSecret),
		jwtRefreshExpiry: getFromEnv("JWT_REFRESH_EXPIRY", defaultRefreshJWTExpiry),
		dbHost:           getFromEnv("DB_HOST", defaultDBHost),
		dbPassword:       getFromEnv("DB_PASSWORD", defaultDBPassword),
		dbUser:           getFromEnv("DB_USERNAME", defaultDBUser),
		dbPort:           getFromEnv("DB_PORT", defaultDBPort),
		dbName:           getFromEnv("DB_DATABASE", defaultDBDatabase),
	}
}

func (env *env) GetJWTSecret() string {
	return env.jwtSecret
}

func (env *env) GetJWTExpiry() string {
	return env.jwtExpiry
}

func (env *env) GetJWTRefreshSecret() string {
	return env.jwtRefreshSecret
}

func (env *env) GetJWTRefreshExpiry() string {
	return env.jwtExpiry
}

func (env *env) GetDbHost() string {
	return env.dbHost
}

func (env *env) GetDbPassword() string {
	return env.dbPassword
}

func (env *env) GetDbUser() string {
	return env.dbUser
}

func (env *env) GetDbPort() string {
	return env.dbPort
}

func (env *env) GetDbName() string {
	return env.dbName
}

var instance *env

func GetInstance() *env {
	if instance == nil {
		instance = newEnv()
	}
	return instance
}

func loadEnv() {
	err := godotenv.Load()
	if err == nil {
		return
	}

	execPath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	envPath := filepath.Join(filepath.Dir(execPath), ".env")
	godotenv.Load(envPath)
}
