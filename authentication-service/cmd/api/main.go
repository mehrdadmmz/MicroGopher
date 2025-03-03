package main

import (
	"authentication/data"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const webPort = "80" // Port for the web server

var counts int64

type Config struct {
	DB     *sql.DB     // Database connection
	Models data.Models // Models for the database tables
}

func main() {
	log.Println("Starting authentication service")

	// connect to DB
	conn := connectToDB()
	if conn == nil {
		log.Panic("Can't connect to Postgres!")
	}

	// set up config
	app := Config{
		DB:     conn,
		Models: data.New(conn),
	}

	// set up routes. This is a method on the Config struct that returns an http.Handler
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

func openDB(dsn string) (*sql.DB, error) {
	// Open the database connection, passing pgx as the driver name, and the DSN which is the connection string
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	// Ping the database to make sure the connection is working. Pinging means sending a request to the database
	// to see if it is still connected. If the database is not connected, the Ping method will return an error.
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

// since we will be adding Postgres to our dockercompose file,  and we need to make sure this is available before we actually return
// database connection, bcz this service might start up before the database does
func connectToDB() *sql.DB {
	dsn := os.Getenv("DSN") // DSN is the data source name, we will get this from the environment variable

	// Keep trying to connect to the database until it is successful
	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("Postgres not ready yet ...")
			counts++
		} else {
			log.Println("Connected to Postgres!")
			return connection
		}

		if counts > 10 {
			log.Println(err)
			return nil
		}

		log.Println("Backing off for 2 seconds....") // overall we will wait for 20 seconds: 10 time * 2 seconds
		time.Sleep(2 * time.Second)
		continue
	}
}
