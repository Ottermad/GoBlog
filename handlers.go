package main

import (
	"database/sql"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/unrolled/render.v1"
	"net/http"
)

func HomeHandler(db *sql.DB, renderer *render.Render) httprouter.Handle {
	return func(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
		posts := getAllPosts(db)
		renderer.HTML(w, http.StatusOK, "index", posts)
	}
}
