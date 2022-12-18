package main

import (
	"authentication/data"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
)

const webPort = "85"

var count int64

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {

	log.Println("Starting authentication service")

	//TODO connect to DB

	//setup config
	app := Config{}

	//setup a webserver
	srv := &http.Server{
		Addr:    fmt.Sprintf("%:s", webPort),
		Handler: app.routes(),
	}

	//start the web server
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}

func OpenDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

//connectToDB maintains connection to the post database 
func connectToDB() *sql.DB {
	dsn := os.Getenv("DSN")

	for {
		connection, err := OpenDB(dsn)
		if err != nil {
			log.Println("Postgres not ready...")
		} else {
			log.Println("Connected to Postgres!")
			return connection
		}
	}
}
