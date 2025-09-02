package main

import (
	"log"

	"github.com/alimohammadi/golan-social.git/internal/env"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8081"),
	}

	app := &application{
		config: cfg,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
