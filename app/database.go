package app

import (
	"database/sql"
	"simple-go-restful/helper"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:T01b01t01@tcp(8.215.72.84:3306)/golang")
	helper.PanicIfError(err)
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)
	return db
}
