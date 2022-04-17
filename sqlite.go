package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// init opens a connection to sqlite and blog.db,
// then set the `db` global variable.
func init() {
	sqlite, err := sql.Open("sqlite3", "blog.db")
	if err != nil {
		log.Println(err)
		return
	}
	db = sqlite
}
