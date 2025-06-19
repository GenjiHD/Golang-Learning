package main

import (
	"log"
	"web-fundamentals/internal/db"
	"web-fundamentals/internal/env"
	"web-fundamentals/internal/store"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://user:adminpassword@localhost/Golang-Web?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONSS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONSS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15min"),
		},
	}

	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)
	if err != nil {
		log.Panic(err)
	}

	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
