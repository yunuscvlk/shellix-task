package main

import (
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"github.com/yunuscvlk/shellix-task/internal/db"
	"github.com/yunuscvlk/shellix-task/internal/router"
)

func main() {
	d := db.Initialize()
	
	r := router.Initialize(d)

	log.Println("Server: started on http://localhost:3000")

	http.ListenAndServe(":3000", r.Mux)
}