package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() *sql.DB {
	//FOR DATABASE LOCALHOST
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/pustaka-api")

	if err != nil {
		fmt.Println(err.Error())
	}

	return db
}
