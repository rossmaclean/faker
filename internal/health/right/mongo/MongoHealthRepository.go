package healthrightmongo

import (
	"context"
	"faker/internal"
	healthcoremodel "faker/internal/health/core/model"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var collection mongo.Collection

type MongoHealthRepository struct {
	collection mongo.Collection
}

func (r *MongoHealthRepository) InitDb() {
	collectionName := "health"

	log.Printf("Connecting to MongoDB %s", collectionName)
	p := internal.GetMongoProperties()

	uri := fmt.Sprintf("mongodb://%s:%s@%s:27017/?authSource=%s",
		p.User, p.Password, p.Host, p.Database)

	clientOptions := options.Client().ApplyURI(uri)
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
	collection = *clnt.Database(p.Database).Collection(collectionName)
}

func (r *MongoHealthRepository) SaveHealth(health healthcoremodel.HealthResponse) error {
	startTime := time.Now()
	_, err := collection.InsertOne(context.TODO(), health)
	if err != nil {
		return err
	}
	runTime := time.Since(startTime)
	log.Printf("Mongo took %s", runTime)
	return nil
}