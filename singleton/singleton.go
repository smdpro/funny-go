package singleton

import (
	"sync"
)

type dbConnection struct {
	ConnectionURI string
}

var instance *dbConnection
var once sync.Once

func GetInstance() *dbConnection {
	once.Do(func() {
		instance = &dbConnection{
			ConnectionURI: "mongodb://localhost:27017",
		}
	})
	return instance
}
