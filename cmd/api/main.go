package main

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

func main() {
	// load .env library
	err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

	// config struct
	cfg := config{
		addr: os.Getenv("ADDR"),
	}
	
	// application struct
	app := &application{
		config: cfg,
	}

	// mount the mux
	mux := app.mount()
	
	// run the app
	log.Fatal(app.run(mux))
}