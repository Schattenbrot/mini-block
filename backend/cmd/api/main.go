package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Schattenbrot/mini-blog/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
}

type AppStatus struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

type application struct {
	config config
	logger *log.Logger
	models models.Models
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "Server port to listen on.")
	flag.StringVar(&cfg.env, "env", "development", "Application environment (development | production)")
	flag.StringVar(&cfg.db.dsn, "dsn", "mongodb://mini-blog-db:27017", "Mongodb dsn to connect to.")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	client, err := openDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}
	db := client.Database("mini-blog")

	app := &application{
		config: cfg,
		logger: logger,
		models: models.NewModels(db),
	}

	serve := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Println("Starting server on port", cfg.port)

	err = serve.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}

func openDB(cfg config) (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(cfg.db.dsn))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return client, err
}