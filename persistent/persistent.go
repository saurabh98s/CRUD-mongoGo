package persistent

import (
	"fmt"
	"log"
	"mongo-golang/config"
	"sync"

	"github.com/globalsign/mgo"
)

type instance struct {
	db *mgo.Database
}

var singleton = &instance{}
var once sync.Once

// Init intialise the mongodb config
func Init() {
	var mongoConfig config.Mongo

	once.Do(func() {
		fmt.Println("connecting to mongo db")
		session, err := mgo.Dial("mongodb://127.0.0.1:27017/")
		if err != nil {
			log.Fatalln(err)
		}
		singleton.db = session.DB(mongoConfig.Database())

		fmt.Println("Connected to mongodb successfully...")
	})
}

// DB returns the database object
func DB() *mgo.Database {
	return singleton.db
}

// Destroy closes the connections & cleans up the instance
func Destroy() error {
	singleton.db.Session.Close()
	return nil
}
