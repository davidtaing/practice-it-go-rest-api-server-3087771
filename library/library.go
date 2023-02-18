package library

import (
	"fmt"
	"log"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!\n")
}

func StartServer() {
	http.HandleFunc("/", helloWorld)
	fmt.Println("Server is running on port 9003.")
	log.Fatal(http.ListenAndServe(":9003", nil))
}