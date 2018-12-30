package models

import (
	"errors"

	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/rogeralbinoi/golang-lab-api/person/database"
	"gopkg.in/go-playground/validator.v9"
)

// Represents a person
type Person struct {
	FirstName string `json:"first_name,omitempty" validate:"required"`
	LastName  string `json:"last_name,omitempty" validate:"required"`
}

// Persists a person
func (r Person) Save() (*mongo.InsertOneResult, error) {
	var validate *validator.Validate

	// init database
	db, ctx := database.Exec()

	// validate fields
	validate = validator.New()
	errs := validate.Struct(r)

	if errs != nil {
		return nil, errors.New("Invalid Person")
	} else {
		return db.Collection("people").InsertOne(ctx, r)
	}
}
