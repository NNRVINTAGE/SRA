package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Access keys
var validAccessKeys = map[string]bool{
	"your-access-key-1": true,
	"your-access-key-2": true,
}

// Middleware to check for access keys
func AccessChecking(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessKey := r.Header.Get("X-Access-Key")
		if validAccessKeys[accessKey] {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}

func main() {
	router := mux.NewRouter()

	// Secure Access Point with Safety Middleware
	router.Use(AccessChecking)

	// routes
	router.HandleFunc("/api/reports", getReports).Methods("GET")
	router.HandleFunc("/api/reports", createReport).Methods("POST")
	router.HandleFunc("/api/reports/{id}", updateReport).Methods("PUT")
	router.HandleFunc("/api/reports/{id}", deleteReport).Methods("DELETE")

	// CORS Configuration
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Update this to the domain of your PHP frontend
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "X-Access-Key"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server started on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
