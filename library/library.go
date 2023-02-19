package library

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type App struct {
	DB 		*sql.DB
	Port	string
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!\n")
}

func (a *App) Initialize() {
	DB, err := sql.Open("sqlite3", "../practiceit.db")
	a.DB = DB

	if err != nil {
		log.Fatal(err.Error())
	}
}

func (a *App) Run() {
	http.HandleFunc("/", helloWorld)
	fmt.Println("Server is running on port ", a.Port)
	log.Fatal(http.ListenAndServe(a.Port, nil))
}