package app

import (
	"database/sql"
	"simple-go-restful/helper"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "")
	helper.PanicIfError(err)
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)
	return db
}
