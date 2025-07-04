package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Environment struct {
  DBUser                  string
  DBPass                  string
  DBHost                  string
  DBPort                  string
  DBSchema                string
  DBMaxIdleConns          string
  DBMaxOpenConns          string
  AppEnv                  string
  AppHost                 string
  AppPort                 string
  AppCorsOrigins          string
  AppCorsHeaders          string
  AppCorsMethods          string
  JWTSecret               string
  JWTAccessTokenExpires   string
  JWTRefreshTokenExpires  string
}

var EnvVar *Environment = new(Environment)

func LoadEnv() {
  err := godotenv.Load(".env")

  if err != nil {
    log.Fatalf("An error occured while loading environment variable: %s", err)
  }

  EnvVar.DBUser = os.Getenv("DB_USER")
  EnvVar.DBPass = os.Getenv("DB_PASS")
  EnvVar.DBHost = os.Getenv("DB_HOST")
  EnvVar.DBPort = os.Getenv("DB_PORT")
  EnvVar.DBSchema = os.Getenv("DB_SCHEMA")
  EnvVar.DBMaxIdleConns = os.Getenv("DB_MAX_IDLE_CONNS")
  EnvVar.DBMaxOpenConns = os.Getenv("DB_MAX_OPEN_CONNS")
  EnvVar.AppEnv = os.Getenv("APP_ENV")
  EnvVar.AppHost = os.Getenv("APP_HOST")
  EnvVar.AppPort = os.Getenv("APP_PORT")
  EnvVar.AppCorsOrigins = os.Getenv("APP_CORS_ORIGINS")
  EnvVar.AppCorsHeaders = os.Getenv("APP_CORS_HEADERS")
  EnvVar.AppCorsMethods = os.Getenv("APP_CORS_METHODS")
  EnvVar.JWTAccessTokenExpires = os.Getenv("JWT_ACCESS_TOKEN_EXPIRES")
  EnvVar.JWTRefreshTokenExpires = os.Getenv("JWT_REFRESH_TOKEN_EXPIRES")
  EnvVar.JWTSecret = os.Getenv("JWT_SECRET")
}