package main

import (
	"database/sql"
	//"fmt"
	"log"
)

// Author struct
type Author struct {
	Forename string
	Surname  string
	table    string
	db       *sql.DB
}

// Author methods
func (author *Author) insertAuthor() (bool, string) {
	query := "INSERT INTO " + author.table + "(FORENAME, SURNAME) VALUES (?, ?)"
	_, err := author.db.Exec(query, author.Forename, author.Surname)

	if err != nil {
		log.Println(err.Error())
		return false, err.Error()
	}

	return true, ""
}

// Author related consts
const authorTable string = "TBL_AUTHOR"

// Author related functions
func createAuthorTable(db *sql.DB) (bool, string) {
	query := "CREATE TABLE IF NOT EXISTS " + authorTable + "(ID INTEGER PRIMARY KEY AUTOINCREMENT, FORENAME VARCHAR(100), SURNAME VACHAR(100))"
	_, err := db.Exec(query)
	if err != nil {
		log.Println(err.Error())
		return false, err.Error()
	}

	query = "CREATE UNIQUE INDEX IF NOT EXISTS INDEX_AUTHORNAME ON " + authorTable + " (FORENAME, SURNAME)"
	_, err = db.Exec(query)
	if err != nil {
		log.Println(err.Error())
		return false, err.Error()
	}

	return true, ""
}
