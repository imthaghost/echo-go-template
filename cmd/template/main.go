package main

import (
	"github.com/imthaghost/echo-template/internal/api"
	"log"
	"math/rand"
	"time"

	"github.com/imthaghost/echo-template/internal/datastore"
)

var startup time.Time

func main() {
	// TODO: keep main clean lets create a setup function that will have all of the inits
	startup = time.Now()
	rand.Seed(time.Now().UTC().UnixNano())
	// TODO: create a config service allow different configs to be passed to other services

	db := &datastore.SQLXService{
		// Configs go in here
	}
	// init database connection
	db.Connect()

	// inject server with needed services
	server := api.NewRESTServer(db)

	log.Printf("starting http server took %v", time.Since(startup))
	// start server
	server.Start()
}