package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/thegera4/go-social-api/internal/store"
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
	
	store := store.NewStorage(nil)

	// application struct
	app := &application{
		config: cfg,
		store: store,
	}

	// mount the mux
	mux := app.mount()
	
	// run the app
	log.Fatal(app.run(mux))
}