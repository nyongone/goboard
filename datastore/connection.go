package datastore

import (
	"database/sql"
	"fmt"
	"go-board-api/config"

	_ "github.com/go-sql-driver/mysql"
)

func OpenConnection() (*sql.DB, error) {
  dsn := GetDSN()

  return sql.Open("mysql", dsn)
}

func GetDSN() string {
  return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", config.EnvVar.DBUser, config.EnvVar.DBPass, config.EnvVar.DBHost, config.EnvVar.DBPort, config.EnvVar.DBSchema)
}