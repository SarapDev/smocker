package models

import (
	"fmt"
	"git.chocodev.kz/rahmet/tools/mocker/internal/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MockData struct {
	Header  map[string]string      `bson:"header" json:"header"`
	Body    map[string]interface{} `bson:"body" json:"body"`
	Timeout int64                  `bson:"timeout" json:"timeout"`
}

func (md MockData) Save() (*mongo.InsertOneResult, error) {
	rep, err := database.RepositoryInstance.Save(md)

	if err != nil {
		return nil, err
	}

	return rep, nil
}

func (md MockData) GetById(id string) database.Model {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println("Invalid id")
	}

	database.RepositoryInstance.GetDataByObjectID(objectId, &md)

	return md
}

func (md MockData) GetCollection() string {
	return "response_model"
}
