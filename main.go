package main

import (
	"context"
	"fmt"
	"git.chocodev.kz/rahmet/tools/mocker/internal"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"net/http"
)

const mongoUrl = "mongodb://root:root@mongodb:27017/?maxPoolSize=20&w=majority"

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoUrl))

	if err != nil {
		panic(err)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected and pinged.")

	internal.NewInstance(client)

	fmt.Println("NewInstance created")

	err = http.ListenAndServe(":8082", nil)

	if err != nil {
		fmt.Println("Server down")
		fmt.Println(err.Error())
		return
	}
}
