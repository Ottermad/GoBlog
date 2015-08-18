package main

import (
	"database/sql"
	//"fmt"
	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/unrolled/render.v1"
	"net/http"
)

// Create db
func NewDB() *sql.DB {
	db, err := sql.Open("sqlite3", "sqlite.db")
	if err != nil {
		panic(err)
	}
	return db
}

func main() {
	db := NewDB()
	createAuthorTable(db)
	createPostTable(db)
	//a := Author{"Derek", "Landy", authorTable, db}
	//a.insertAuthor()
	//p := Post{a, "Hello Again", "This is my second post", postTable}
	//p.insertPost()
	//fmt.Println(getAllPosts(db))
	renderer := render.New(render.Options{})

	router := httprouter.New()

	// Home Route
	router.GET("/", HomeHandler(db, renderer))
	http.ListenAndServe(":8080", router)
}
