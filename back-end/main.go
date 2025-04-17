package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/Oustaa/go-react-backend/routes"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8000, "the port where the app will listen on")
	flag.Parse()

	addr := fmt.Sprintf(":%d", port)

	r := routes.NewRouters()

	server := http.Server{
		Addr:    addr,
		Handler: r,
	}

	log.Printf("The app is listening on port %d", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
