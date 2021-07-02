package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"waterjugserver/internal/server"
)

func main() {
	//Starting server
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { server.Handler(w, r) }).Methods("POST")
	fmt.Println("starting server at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
