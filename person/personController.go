package main

import (
	"encoding/json"
	"net/http"

	"github.com/rogeralbinoi/golang-lab-api/person/models"
)

func create(w http.ResponseWriter, r *http.Request) {
	firstName, lastName := r.FormValue("first_name"), r.FormValue("last_name")
	person := models.Person{FirstName: firstName, LastName: lastName}

	document, err := person.Save()

	if err != nil {
		http.Error(w, "Invalid", 400)
	}

	json.NewEncoder(w).Encode(document)
}
