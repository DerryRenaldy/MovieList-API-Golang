package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"movie_api/helper"
)

func NewDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/cinema_booking2?parseTime=true")
	helper.PrintError(err)

	err = db.Ping()
	if err != nil {
		log.Error(err)
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("Server Listening On Localhost:8010")
	return db, nil

}
