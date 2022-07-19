package internal

import (
	"fmt"
	"github.com/magiconair/properties"
	"os"
)

var p *properties.Properties

//func init() {
//	filename := "./configs/"
//	log.Println("Initialising config")
//	switch GetEnv() {
//	case "prod":
//		filename += "prod.properties"
//		break
//	default:
//		err := os.Setenv("ENV", "local")
//		if err != nil {
//			log.Fatal(err)
//		}
//		filename += "local.properties"
//		break
//	}
//	p = properties.MustLoadFile(filename, properties.UTF8)
//	log.Printf("Running with %s profile", GetEnv())
//}

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
