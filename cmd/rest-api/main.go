package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/keshav-sudo/rest-api/internal/config"
)

func main() {
	// Load config
	cfg := config.MustLoad()

	// Create router
	router := http.NewServeMux()

	// Health check endpoint
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("server is healthy"))
	})

	// Server setup
	server := &http.Server{
		Addr:    cfg.HTTPServer.Addr,
		Handler: router,
	}

	fmt.Println("Server starting at", cfg.HTTPServer.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
