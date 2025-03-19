package config

import "os"

const mongoConnection string = "MONGO_CONNECTION"

func StringConnectionMongo() string {
	if len(os.Getenv(mongoConnection)) > 0 {
		return os.Getenv(mongoConnection)
	}
	return "mongodb://localhost:27017"
}
