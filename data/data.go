package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

var Db *sql.DB

func init() {
	var err error
	dbCredentials := fmt.Sprintf("dbname=%s dbuser=%s dbpass=%s sslmode=disable", config.DbName, config.DbUser, config.DbPassword)
	_, err = sql.Open("postgres", dbCredentials)
	if err != nil {
		log.Fatal("Could not connect to DB ", err)
	}
	return
}
