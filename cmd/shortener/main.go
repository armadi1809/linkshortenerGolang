package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"

	"azizrmadi.net.shortenUrls/internal/models"
	_ "github.com/mattn/go-sqlite3"
)

type application struct {
	urls *models.ShortUrlModel
}

func main() {
	connectionString := flag.String("dbConnString", "./urlMap", "sqlite connection string to store the mappings")

	db, err := openDb(*connectionString)

	app := &application{urls: &models.ShortUrlModel{DB: db}}

	srv := &http.Server{
		Addr:    ":8080",
		Handler: app.routes(),
	}
	if err != nil {
		log.Fatal(err)
	}

	err = srv.ListenAndServe()
	log.Fatal(err)

}

func openDb(connectionString string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", connectionString)

	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}

	return db, nil
}
