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

// Host returns the Mongo host
func (mongoConfig *Mongo) Host() string {
	return mongoConfig.host
}

// Port returns the mongo Port
func (mongoConfig *Mongo) Port() string {
	return mongoConfig.port
}

// RootPassword returns the mongo RootPassword
func (mongoConfig *Mongo) RootPassword() string {
	return mongoConfig.rootPassword
}

// User returns the mongo user
func (mongoConfig *Mongo) User() string {
	return mongoConfig.user
}

// UserPassword returns the mongo userPassword
func (mongoConfig *Mongo) UserPassword() string {
	return mongoConfig.userPassword
}

// Database returns the mongo Database
func (mongoConfig *Mongo) Database() string {
	return mongoConfig.database
}
// ConnectionString returns a connection string for mongoDial
func (mongoConfig *Mongo) ConnectionString() string {
	return fmt.Sprintf("mongodb://%v:%v@%v:%v/%v", mongoConfig.user, mongoConfig.userPassword, mongoConfig.host, mongoConfig.port, mongoConfig.database)
}
