package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "80" // the port the web server will listen on, we will be using port 80 since docker will map it to 8080

type Config struct{}

func main() {
	app := Config{}

	log.Printf("Starting broker service on port %s\n", webPort)

	// define http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort), // listen on port 80
		Handler: app.routes(),                // use the routes defined in routes.go
	}

	// start the server
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
