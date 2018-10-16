package main

import (
	"encoding/json"
	"net/http"
)

var persons = []Person{
	Person{FirstName: "Will", LastName: "Byers"},
	Person{FirstName: "Jonathan", LastName: "Byers"},
}

func getAll(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(persons)
}
