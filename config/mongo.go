package config

import (
	"fmt"
	"os"

	"github.com/subosito/gotenv"
)

// Mongo is the config for the db vars
type Mongo struct {
	host         string
	port         string
	rootPassword string
	user         string
	userPassword string
	database     string
}

// Load loads the config for the mongodb
func (mongoConfig *Mongo) Load() {
	fmt.Println("Reading mongo config...")
	err := gotenv.Load()
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	mongoConfig.host = os.Getenv("MONGODB_HOST")
	mongoConfig.port = os.Getenv("MONGODB_PORT")
	mongoConfig.rootPassword = os.Getenv("MONGODB_ROOT_PASSWORD")
	mongoConfig.user = os.Getenv("MONGODB_USER")
	mongoConfig.userPassword = os.Getenv("MONGODB_PASSWORD")
	mongoConfig.database = os.Getenv("MONGODB_DATABASE")
	fmt.Println("Read config..all okay")
}
