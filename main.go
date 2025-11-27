package main

import (
	"github.com/sturrdhq/celery-server/config"
	"github.com/sturrdhq/celery-server/internal/api"
	"github.com/sturrdhq/celery-server/internal/database"
	"log"
)

func main() {
	config.SetupLogger()

	// set up a dabase client
	dbClient, err := database.NewDBClient()
	if err != nil {
		log.Fatal("failed to create database client", err)
	}

	// get the client database connection
	conn, err := dbClient.DB.DB()
	if err != nil {
		log.Fatal("failed to get database connection", err)
	}

	defer func() {
		log.Println("closing database connection")
		_ = conn.Close()
		log.Println("database connection closed")
	}()

	// run the setup required to get the database ready
	// migrations and other setup
	err = dbClient.Setup()
	if err != nil {
		log.Fatal("failed to setup database", err)
	}

	PORT := 8000
	log.Printf("starting server on port %d", PORT)
	api.InitServer(PORT, dbClient)
}
