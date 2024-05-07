package main

import (
	"net/http"

	"github.com/rs/cors"
)

func main() {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://dmm.com", "http://localhost:3000"},
		AllowedMethods:   []string{http.MethodGet},
		AllowCredentials: true,
		Debug:            true,
	})

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"body\": \"Hello, World!\"}"))
	})

	http.ListenAndServe(":8000", c.Handler(handler))
}
