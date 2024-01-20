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
	defaultJWTExpiry        = 86400
	defaultRefreshJWTSecret = "dummy_refresh"
	defaultRefreshJWTExpiry = 1209600
)

type Env struct {
	jWTToken        string
	jWTRefreshToken string
	dbHost          string
	dbPassword      string
	dbUser          string
	dbPort          string
	dbName          string
}

func getFromEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value != "" {
		return value
	}
	return defaultValue
}

func newEnv() *Env {
	loadEnv()
	return &Env{
		jWTToken:        getFromEnv("JWT_SECRET", defaultJWTSecret),
		jWTRefreshToken: getFromEnv("JWT_REFRESH_SECRET", defaultRefreshJWTSecret),
		dbHost:          getFromEnv("DB_HOST", defaultDBHost),
		dbPassword:      getFromEnv("DB_PASSWORD", defaultDBPassword),
		dbUser:          getFromEnv("DB_USERNAME", defaultDBUser),
		dbPort:          getFromEnv("DB_PORT", defaultDBPort),
		dbName:          getFromEnv("DB_DATABASE", defaultDBDatabase),
	}
}

func (env *Env) GetJWTToken() string {
	return env.jWTToken
}

func (env *Env) GetJWTRefreshToken() string {
	return env.jWTRefreshToken
}

func (env *Env) GetDbHost() string {
	return env.dbHost
}

func (env *Env) GetDbPassword() string {
	return env.dbPassword
}

func (env *Env) GetDbUser() string {
	return env.dbUser
}

func (env *Env) GetDbPort() string {
	return env.dbPort
}

func (env *Env) GetDbName() string {
	return env.dbName
}

var instance *Env

func GetInstance() *Env {
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
