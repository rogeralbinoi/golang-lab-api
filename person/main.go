package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	port := ":3000"

	// routers
	router.HandleFunc("/person", create).Methods("POST")

	// init server
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, router))
}
