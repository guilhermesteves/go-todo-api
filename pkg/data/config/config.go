package config

import (
	"os"
)

func MongoDbDsn() string {
	return os.Getenv("MONGODB_DSN")
}
func MongoDbDatabase() string {
	return os.Getenv("MONGODB_DATABASE")
}
