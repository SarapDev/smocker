package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
)

var once sync.Once

type Repository struct {
	client *mongo.Client
	db     string
}

var RepositoryInstance *Repository

func GetInstance(db string, client *mongo.Client) *Repository {
	if RepositoryInstance == nil {
		once.Do(
			func() {
				fmt.Println("Creating single instance now.")
				RepositoryInstance = &Repository{client, db}
			})
	} else {
		fmt.Println("Single instance already created.")
	}

	return RepositoryInstance
}

func (r Repository) GetDataByObjectID(oId primitive.ObjectID, m Model) {
	coll := r.client.Database(r.db).Collection(m.GetCollection())
	data := bson.M{"_id": oId}

	err := coll.FindOne(context.Background(), data).Decode(m)
	if err != nil {
		return
	}
}

func (r Repository) Save(m Model) (*mongo.InsertOneResult, error) {
	coll := r.client.Database(r.db).Collection(m.GetCollection())

	result, err := coll.InsertOne(context.Background(), m)
	if err != nil {
		return nil, err
	}

	return result, nil
}
