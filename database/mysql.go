package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/joho/godotenv"
)

func NewMySQL() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		panic("can't get environment variables")
	}

	mysqlName := os.Getenv("MYSQL_NAME")
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")
	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlPort := os.Getenv("MYSQL_PORT")

	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", mysqlName, mysqlPassword, mysqlHost, mysqlPort, mysqlDatabase)

	DB, err := sql.Open("mysql", dns)

	if err != nil {
		panic("can't connect to MySQL database")
	}

	DB.SetMaxIdleConns(10)
	DB.SetMaxOpenConns(100)
	DB.SetConnMaxIdleTime(5 * time.Minute)
	DB.SetConnMaxLifetime(60 * time.Minute)

	return DB
}
