package internal

import (
	"git.chocodev.kz/rahmet/tools/mocker/internal/database"
	"git.chocodev.kz/rahmet/tools/mocker/internal/network"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

const dbName = "mocker"

func NewInstance(client *mongo.Client) {
	go database.GetInstance(dbName, client)

	http.HandleFunc("/generate", network.GenerateUrlWithData)

	http.HandleFunc("/mocker/", network.GetMockData)
}
