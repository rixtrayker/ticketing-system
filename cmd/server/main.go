package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rixtrayker/ticketing-system/internal/db"
	"github.com/rixtrayker/ticketing-system/internal/graph"
)

const defaultPort = "8080"

func main() {
	// Initialize database connection
	dbConfig := db.NewConfig()
	if err := db.Connect(dbConfig); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Run migrations
	if err := db.AutoMigrate(); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Create GraphQL resolver with database connection
	resolver := &graph.Resolver{
		DB: db.DB,
	}

	// We'll uncomment this after generating the GraphQL files
	// srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
	// 	Resolvers: resolver,
	// }))

	// Add middleware for CORS and other headers
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	// http.Handle("/query", corsMiddleware(srv))

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	log.Printf("Server is running on http://localhost:%s", port)
	log.Printf("GraphQL Playground available at http://localhost:%s/", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// corsMiddleware adds CORS headers to the response
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
} 