package internal

import (
	"fmt"
	"github.com/magiconair/properties"
	"log"
	"os"
	"strconv"
)

var p *properties.Properties

func init() {
	log.Println("Initialising config")

	filename := "./configs/"

	if len(os.Args) == 2 {
		filename = os.Args[1]
	}

	switch GetEnv() {
	case "prod":
		filename += "prod.properties"
		break
	default:
		err := os.Setenv("ENV", "local")
		if err != nil {
			log.Fatal(err)
		}
		filename += "local.properties"
		break
	}

	p = properties.MustLoadFile(filename, properties.UTF8)
	log.Printf("Running with %s profile", GetEnv())
}

func GetEnv() string {
	return os.Getenv("ENV")
}

func IsLocal() bool {
	return GetEnv() == "local"
}

type MongoProperties struct {
	ConnectionString string
	Database         string
}

func GetMongoProperties() MongoProperties {
	mongoHost := p.MustGetString("mongo.host")
	mongoDatabase := p.MustGetString("mongo.database")
	mongoUser := p.MustGetString("mongo.user")
	mongoPassword := p.GetString("mongo.password", os.Getenv("MONGO_PASSWORD"))
	connectionString := ""
	if IsLocal() {
		connectionString = fmt.Sprintf("mongodb://%s:27017/?authSource=%s", mongoHost, mongoDatabase)
	} else {
		connectionString = fmt.Sprintf("mongodb://%s:%s@%s:27017/?authSource=%s",
			mongoUser, mongoPassword, mongoHost, mongoDatabase)
	}
	return MongoProperties{
		ConnectionString: connectionString,
		Database:         mongoDatabase,
	}
}

type FakerProperties struct {
	NumToGenerate int
	SleepMillis   int
}

func GetFakerProperties() FakerProperties {
	numToGenerate := p.GetString("faker.num-to-generate", os.Getenv("FAKER_NUM_TO_GENERATE"))
	sleepMillis := p.GetString("faker.sleep-millis", os.Getenv("FAKER_SLEEP_MILLIS"))
	return FakerProperties{
		NumToGenerate: convertStringToInt64(numToGenerate),
		SleepMillis:   convertStringToInt64(sleepMillis),
	}
}

func convertStringToInt64(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return num
}
