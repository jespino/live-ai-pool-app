package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jespino/pool-app/internal/database"
	"github.com/jespino/pool-app/internal/handlers"
)

func main() {
	var storage database.Storage

	// Determine which database to use
	dbType := os.Getenv("DB_TYPE")
	if dbType == "postgres" {
		// Get connection details from environment
		connStr := os.Getenv("DB_CONNECTION_STRING")
		if connStr == "" {
			// Provide a default connection string if not specified
			host := getEnv("DB_HOST", "localhost")
			port := getEnv("DB_PORT", "5432")
			user := getEnv("DB_USER", "postgres")
			password := getEnv("DB_PASSWORD", "postgres")
			dbname := getEnv("DB_NAME", "poolapp")
			sslmode := getEnv("DB_SSLMODE", "disable")

			connStr = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
				host, port, user, password, dbname, sslmode)
		}

		// Connect to PostgreSQL
		pgDB, err := database.NewPostgresDB(connStr)
		if err != nil {
			log.Fatalf("Failed to connect to PostgreSQL: %v", err)
		}

		// Initialize the database schema
		if err := pgDB.Initialize(); err != nil {
			log.Fatalf("Failed to initialize PostgreSQL schema: %v", err)
		}

		storage = pgDB
		log.Println("Using PostgreSQL database")
	} else {
		// Use in-memory database by default
		storage = database.NewMemoryDB()
		log.Println("Using in-memory database")
	}

	// Set up the handler
	h := handlers.NewHandler(storage)

	// Create a new router
	r := mux.NewRouter()

	// Define API routes
	apiRouter := r.PathPrefix("/api").Subrouter()

	// Pool routes
	apiRouter.HandleFunc("/pools", h.CreatePool).Methods("POST")
	apiRouter.HandleFunc("/pools", h.ListPools).Methods("GET")
	apiRouter.HandleFunc("/pools/{id}", h.GetPool).Methods("GET")
	apiRouter.HandleFunc("/pools/{id}/vote", h.Vote).Methods("POST")

	// User routes
	apiRouter.HandleFunc("/users", h.CreateUser).Methods("POST")

	// Set up CORS
	corsMiddleware := gohandlers.CORS(
		gohandlers.AllowedOrigins([]string{"*"}),
		gohandlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		gohandlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	// Determine port to use
	port := 8081
	if portStr := os.Getenv("PORT"); portStr != "" {
		if p, err := strconv.Atoi(portStr); err == nil {
			port = p
		}
	}

	// Start the server
	fmt.Printf("Server starting on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), corsMiddleware(r)))
}

// Helper function to get environment variables with defaults
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

