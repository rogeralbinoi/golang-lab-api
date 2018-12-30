package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rogeralbinoi/golang-lab-api/person/controllers"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	port := ":3000"

	// routers
	router.HandleFunc("/person", personController.Create).Methods("POST")

	// init server
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, router))
}
