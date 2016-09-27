package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/plispe/teamguru/handler"
)

func main() {
	// Get port from ENV
	port := os.Getenv("PORT")

	// Asign default port
	if port == "" {
		port = "80"
	}

	// Log info
	log.Println(fmt.Sprintf(`Server listening on port: "%s"`, port))
	// Handle request
	http.ListenAndServe(":"+port, handler.NewHandler())
}
