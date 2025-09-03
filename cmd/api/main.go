package main

import (
	"log"

	"github.com/alimohammadi/golan-social.git/internal/env"
	"github.com/alimohammadi/golan-social.git/internal/store"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8081"),
	}
	store := store.NewStorage(nil)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
