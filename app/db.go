package app

import (
	"database/sql"
	"fmt"
	"go_blogger/helper"
	"os"
	"time"
)

func NewDB() *sql.DB {
	//db_env
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

func Dbtest() *sql.DB {
	db, err := sql.Open("postgres", "user=postgres password=ifqygazhar dbname=go_blogger port=5432 sslmode=disable")
	helper.PanicIfError(err)

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
