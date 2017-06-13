package controllers

import (
	"database/sql"
	"fmt"
     _ "github.com/mattn/go-sqlite3"
	"github.com/revel/revel"
)

//DB represents the database instance
var db *sql.DB

//InitDB initializes DB connection
func InitDB() {
	connstring := fmt.Sprintf("/Users/programmer/go/src/SnowGo/books.db")

	var err error
	db, err = sql.Open("sqlite3", connstring)
	if err != nil {
		revel.INFO.Println("DB Error", err)
	}
	revel.INFO.Println("DB Connected")
}

func init() {
	revel.OnAppStart(InitDB)
}
