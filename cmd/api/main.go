package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Blue-Onion/go-monolith/internal/config"
)

func main() {
	cfg := config.MustLoad()
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"OK"}`))
	})
	srv := http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
		IdleTimeout:  time.Second * 10,
	}
	log.Printf("http://localhost:%s", cfg.Port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("server failed %v", err)
	}
}
