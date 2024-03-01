package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tttuepbergen/tsis2/pkg/abr-plus/model"

	_ "github.com/lib/pq"
)

type config struct {
	port string
	env  string
	db   struct {
		dsn string
	}
}

type application struct {
	config config
	models model.Models
}

func main() {
	var cfg config
	flag.StringVar(&cfg.port, "port", ":8081", "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&cfg.db.dsn, "db-dsn", "postgres://codev0:pa55word@localhost/lecture6?sslmode=disable", "PostgreSQL DSN")
	flag.Parse()

	// Connect to DB
	db, err := openDB(cfg)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	app := &application{
		config: cfg,
		models: model.NewModels(db),
	}

	app.run()
}

func (app *application) run() {
	r := mux.NewRouter()

	v1 := r.PathPrefix("/api/v1").Subrouter()

	// Animal Singleton
	// Create a new animal
	v1.HandleFunc("/animals", app.createAnimalHandler).Methods("POST")
	// Get a specific animal
	v1.HandleFunc("/animals/{animalId:[0-9]+}", app.getAnimalHandler).Methods("GET")
	// Update a specific animal
	v1.HandleFunc("/animals/{animalId:[0-9]+}", app.updateAnimalHandler).Methods("PUT")
	// Delete a specific animal
	v1.HandleFunc("/animals/{animalId:[0-9]+}", app.deleteAnimalHandler).Methods("DELETE")

	log.Printf("Starting server on %s\n", app.config.port)
	err := http.ListenAndServe(app.config.port, r)
	log.Fatal(err)
}

func openDB(cfg config) (*sql.DB, error) {
	// Use sql.Open() to create an empty connection pool, using the DSN from the config // struct.
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
