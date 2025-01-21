package main

import (
	"log"
	"github.com/joho/godotenv"
	"github.com/thegera4/go-social-api/cmd/api/utils"
	"github.com/thegera4/go-social-api/internal/db"
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
		addr: utils.LoadAddr(),
		db: dbConfig{
			addr: utils.LoadDBAddr(),
			maxOpenConns: utils.LoadMaxOpenConns(),
			maxIdleConns: utils.LoadMaxIdleConns(),
			maxIdleTime: utils.LoadMaxIdleTime(),
		},
	}
	
	// db connection
	db, err := db.New(cfg.db.addr, cfg.db.maxOpenConns, cfg.db.maxIdleConns, cfg.db.maxIdleTime)
	if err != nil { log.Panic(err) }
	defer db.Close()
	log.Println("Database connection pool established")

	store := store.NewStorage(db)

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