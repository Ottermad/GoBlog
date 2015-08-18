package main

import (
	"database/sql"
	//"fmt"
	"log"
)

// Post struct
type Post struct {
	Author  Author
	Title   string
	content string
	table   string
}

// Post methods
func (post *Post) insertPost() (bool, string) {
	// Get Author id
	var id int
	post.Author.db.QueryRow("SELECT ID FROM TBL_AUTHOR WHERE FORENAME = ? AND SURNAME = ?", post.Author.Forename, post.Author.Surname).Scan(&id)

	_, err := post.Author.db.Exec("INSERT INTO TBL_POST(AUTHOR_ID, TITLE, CONTENT) VALUES(?, ?, ?)", id, post.Title, post.content)

	if err != nil {
		log.Println(err.Error())
		return false, err.Error()
	}
	return true, ""
}

// Post related consts/vars
const postTable string = "TBL_POST"

// Post related functions
func createPostTable(db *sql.DB) (bool, string) {
	query := "CREATE TABLE IF NOT EXISTS " + postTable + "(ID INTEGER PRIMARY KEY AUTOINCREMENT, AUTHOR_ID INTEGER, TITLE VARCHAR(100), CONTENT TEXT, FOREIGN KEY (AUTHOR_ID) REFERENCES " + authorTable + "(ID))"
	_, err := db.Exec(query)
	if err != nil {
		log.Println(err.Error())
		return false, err.Error()
	}

	query = "CREATE UNIQUE INDEX IF NOT EXISTS INDEX_POSTTITLE ON " + postTable + " (TITLE)"
	_, err = db.Exec(query)
	if err != nil {
		log.Println(err.Error())
		return false, err.Error()
	}

	return true, ""
}

/*func getAllPosts(db *sql.DB) map[Author][]Post {
	rows, err := db.Query("SELECT * FROM " + postTable)

	data := make(map[Author][]Post)

	if err != nil {
		log.Println("ERROR @ SELECT * FROM postTable")
		log.Println(err.Error())
		return data
	}

	defer rows.Close()

	for rows.Next() {
		var id, author_id int
		var title, content string

		if err := rows.Scan(&id, &author_id, &title, &content); err != nil {
			log.Println(err.Error())
			return data
		}

		// Get author
		var forename, surname string
		err = db.QueryRow("SELECT FORENAME, SURNAME FROM  "+authorTable+" WHERE ID = ?", author_id).Scan(&forename, &surname)

		if err != nil {
			log.Println(err.Error())
		}

		a := Author{forename, surname, authorTable, db}
		p := Post{a, title, content, postTable}

		var isKey bool = false

		for k := range data {
			if (k.forename == a.forename) && (k.surname == a.surname) {
				data[k] = append(data[k], p)
			}
		}

		if !isKey {
			data[a] = []Post{p}
		}

	}

	return data

}*/

func getAllPosts(db *sql.DB) []Post {
	rows, err := db.Query("SELECT * FROM " + postTable)

	var data []Post

	if err != nil {
		log.Println("ERROR @ SELECT * FROM postTable")
		log.Println(err.Error())
		return data
	}

	defer rows.Close()

	for rows.Next() {
		var id, author_id int
		var title, content string

		if err := rows.Scan(&id, &author_id, &title, &content); err != nil {
			log.Println(err.Error())
			return data
		}

		// Get author
		var forename, surname string
		err = db.QueryRow("SELECT FORENAME, SURNAME FROM  "+authorTable+" WHERE ID = ?", author_id).Scan(&forename, &surname)

		if err != nil {
			log.Println(err.Error())
		}

		a := Author{forename, surname, authorTable, db}
		p := Post{a, title, content, postTable}
		data = append(data, p)
	}

	return data

}
