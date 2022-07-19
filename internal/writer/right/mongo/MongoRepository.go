package writerrightmongo

import (
	"context"
	"faker/internal"
	"faker/internal/writer/core/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var client mongo.Client
var collection mongo.Collection

type MongoPeopleRepository struct {
	client     mongo.Client
	collection mongo.Collection
}

func (r *MongoPeopleRepository) Ping() error {
	return client.Ping(context.TODO(), nil)
}

func (r *MongoPeopleRepository) InitDb() {
	collectionName := "people"

	log.Printf("Connecting to MongoDB %s", collectionName)
	p := internal.GetMongoProperties()

	clientOptions := options.Client().ApplyURI(p.ConnectionString)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clnt, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Println(err)
	}
	err = clnt.Ping(context.TODO(), nil)

	if err != nil {
		log.Println(err)
	}
	log.Printf("Connected to MongoDB %s", collectionName)
	client = *clnt
	collection = *clnt.Database(p.Database).Collection(collectionName)
}

func (r *MongoPeopleRepository) SavePerson(person writercoremodel.Person) error {
	_, err := collection.InsertOne(context.TODO(), person)
	return err
}
