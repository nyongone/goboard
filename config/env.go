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
  AppEnv                  string
  AppHost                 string
  AppPort                 string
  AppCors                 string
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
  EnvVar.AppEnv = os.Getenv("APP_ENV")
  EnvVar.AppHost = os.Getenv("APP_HOST")
  EnvVar.AppPort = os.Getenv("APP_PORT")
  EnvVar.AppCors = os.Getenv("APP_CORS")
  EnvVar.JWTAccessTokenExpires = os.Getenv("JWT_ACCESS_TOKEN_EXPIRES")
  EnvVar.JWTRefreshTokenExpires = os.Getenv("JWT_REFRESH_TOKEN_EXPIRES")
  EnvVar.JWTSecret = os.Getenv("JWT_SECRET")
}