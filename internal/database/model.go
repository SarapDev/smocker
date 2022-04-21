package database

import "go.mongodb.org/mongo-driver/mongo"

type Model interface {
	Save() (*mongo.InsertOneResult, error)
	GetById(id string) Model
	GetCollection() string
}
