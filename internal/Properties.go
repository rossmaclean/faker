package internal

import (
	"github.com/magiconair/properties"
	"log"
	"os"
)

var p *properties.Properties

func init() {
	filename := "./configs/"
	switch GetEnv() {
	case "prod":
		filename += "prod.properties"
		break
	default:
		os.Setenv("ENV", "local")
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
	User                string
	Password            string
	Host                string
	Database            string
}

func GetMongoProperties() MongoProperties {
	mongoPassword := p.GetString("mongo.password", os.Getenv("MONGO_PASSWORD"))
	return MongoProperties{
		User:                p.MustGetString("mongo.user"),
		Password:            mongoPassword,
		Host:                p.MustGetString("mongo.host"),
		Database:            p.MustGetString("mongo.database"),
	}
}
