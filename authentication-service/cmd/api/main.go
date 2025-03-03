package main

import "database/sql"

const webPort = "80" // Port for the web server

type Config struct {
	DB     *sql.DB     // Database connection
	Models data.Models // Models for the database tables
}

func main() {

}
