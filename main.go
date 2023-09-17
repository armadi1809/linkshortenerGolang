package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	connectionString := flag.String("dbConnString", "./urlMap", "sqlite connection string to store the mappings")

	db, err := openDb(*connectionString)

	if err != nil {
		log.Fatal("Cannot open the database")
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/path/", func(w http.ResponseWriter, r *http.Request) {
		var path string
		if strings.HasPrefix(r.URL.Path, "/path/") {
			path = r.URL.Path[len("/path/"):]
			fmt.Println("Path: " + r.URL.Path)
		}
		fmt.Println("Path: " + r.URL.Path)
		url := dbUrlFinder(path, db)
		fmt.Println("Redirecting to " + url)
		http.Redirect(w, r, url, http.StatusPermanentRedirect)
		return
	})
	fmt.Println("Now listening on port 8080")
	http.ListenAndServe(":8080", mux)

}

func openDb(connectionString string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", connectionString)

	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}

	return db, nil
}

func dbUrlFinder(path string, db *sql.DB) string {
	query := `SELECT path, url FROM urls WHERE path=?`
	var pathFound string
	err := db.QueryRow(query, path).Scan(&path, &pathFound)
	if err != nil {
		log.Fatal(err.Error())
		return ""
	}

	return pathFound

}
