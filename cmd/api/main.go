package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("alll ookk"))
	})
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatalf("server failed %v", err)
	}
}
