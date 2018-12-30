package database

import (
	"context"
	"time"

	"github.com/mongodb/mongo-go-driver/mongo"
)

func Exec() (*mongo.Database, context.Context) {
	client, err := mongo.NewClient("mongodb://localhost:27017")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	if err != nil {
		panic(err)
	}

	err = client.Connect(ctx)

	if err != nil {
		panic(err)
	}

	return client.Database("person"), ctx
}
