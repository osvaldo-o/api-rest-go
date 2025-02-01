package mysql_db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func ConnectClient() (client *sql.DB, err error) {

	err_os := godotenv.Load()
	if err_os != nil {
		fmt.Println("Error cargando el archivo .env")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	databases, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = databases.Ping()
	if err != nil {
		return nil, err
	}

	return databases, nil

}
