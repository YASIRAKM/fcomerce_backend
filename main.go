package main

import (
	"fcomerce/db"
	"fcomerce/router"
	"log"
	"os"
)

func main() {

	db.Init()
	defer db.CloseDB()
	e := router.New()

	port := os.Getenv("PORT")
	if port == "" {
	    log.Fatal("PORT environment variable is not set")
	}

	// Start the server
	e.Logger.Fatal(e.Start(":" + port))
	// e.Logger.Fatal(e.Start(":8000"))

}
