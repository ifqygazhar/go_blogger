package app

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"go_blogger/helper"
	"os"
	"time"
)

func NewDB() *sql.DB {
	//db_env
	godotenv.Load(".env")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")


	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	helper.PanicIfError(err)

	// close database
	defer db.Close()

	// check db
	err = db.Ping()
	helper.PanicIfError(err)

	//set_time_db
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	fmt.Println("Connected!")

	return db
}
