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

	router.HandleFunc("/", getAll).Methods("GET")

	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, router))

}
