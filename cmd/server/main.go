package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rixtrayker/ticketing-system/internal/db"
	"github.com/rixtrayker/ticketing-system/internal/graph"
	"github.com/rixtrayker/ticketing-system/internal/graph/generated"
	"github.com/rixtrayker/ticketing-system/internal/repository"
	"github.com/rixtrayker/ticketing-system/internal/service"
)

const (
	defaultPort           = "8080"
	shutdownTimeout       = 30 * time.Second
	readTimeout          = 15 * time.Second
	writeTimeout         = 15 * time.Second
	idleTimeout          = 60 * time.Second
	readHeaderTimeout    = 5 * time.Second
)

// HealthResponse represents the health check response
type HealthResponse struct {
	Status    string            `json:"status"`
	Timestamp time.Time         `json:"timestamp"`
	Version   string            `json:"version,omitempty"`
	Checks    map[string]string `json:"checks,omitempty"`
}

func main() {
	// Initialize structured logging
	logger := log.New(os.Stdout, "[TICKETING-SYSTEM] ", log.LstdFlags|log.Lshortfile)
	
	// Get configuration
	config := getConfig()
	logger.Printf("Starting server on port %s", config.Port)

	// Initialize database connection
	dbConfig := db.NewConfig()
	if err := db.Connect(dbConfig); err != nil {
		logger.Fatalf("Failed to connect to database: %v", err)
	}
	logger.Println("Database connection established")

	// Run migrations
	if err := db.AutoMigrate(); err != nil {
		logger.Fatalf("Failed to run migrations: %v", err)
	}
	logger.Println("Database migrations completed")

	// Initialize repositories
	ticketRepo := repository.NewTicketRepository(db.DB)
	userRepo := repository.NewUserRepository(db.DB)
	assetRepo := repository.NewAssetRepository(db.DB)

	// Initialize services
	ticketService := service.NewTicketService(ticketRepo, userRepo, assetRepo)

	// Create GraphQL resolver with dependencies
	resolver := &graph.Resolver{
		DB:            db.DB,
		TicketService: ticketService,
	}

	// Create GraphQL server with configuration
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: resolver,
		Directives: generated.DirectiveRoot{
			// Add any custom directives here
		},
		Complexity: generated.ComplexityRoot{
			// Add any custom complexity functions here
		},
	}))

	// GraphQL server is ready (no additional middleware needed)

	// Create HTTP server
	mux := http.NewServeMux()
	
	// Health check endpoint
	mux.HandleFunc("/health", healthCheckHandler(logger))
	
	// Metrics endpoint (basic)
	mux.HandleFunc("/metrics", metricsHandler(logger))
	
	// GraphQL endpoints
	if config.Environment == "development" {
		mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
		logger.Println("GraphQL Playground enabled at /")
	}
	mux.Handle("/query", corsMiddleware(recoveryMiddleware(loggingMiddleware(graphqlMiddleware(logger)(srv), logger))))

	// Configure HTTP server with production settings
	server := &http.Server{
		Addr:              ":" + config.Port,
		Handler:           mux,
		ReadTimeout:       readTimeout,
		WriteTimeout:      writeTimeout,
		IdleTimeout:       idleTimeout,
		ReadHeaderTimeout: readHeaderTimeout,
		ErrorLog:          logger,
	}

	// Channel to listen for interrupt signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Start server in a goroutine
	go func() {
		logger.Printf("Server starting on http://localhost:%s", config.Port)
		if config.Environment == "development" {
			logger.Printf("GraphQL Playground available at http://localhost:%s/", config.Port)
		}
		logger.Printf("GraphQL API available at http://localhost:%s/query", config.Port)
		logger.Printf("Health check available at http://localhost:%s/health", config.Port)
		
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Block until we receive a signal
	<-quit
	logger.Println("Shutting down server...")

	// Create a context with timeout for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	// Attempt graceful shutdown
	if err := server.Shutdown(ctx); err != nil {
		logger.Printf("Server forced to shutdown: %v", err)
		return
	}

	// Close database connection
	if sqlDB, err := db.DB.DB(); err == nil {
		if err := sqlDB.Close(); err != nil {
			logger.Printf("Error closing database connection: %v", err)
		} else {
			logger.Println("Database connection closed")
		}
	}

	logger.Println("Server gracefully stopped")
}

// Config holds application configuration
type Config struct {
	Port        string
	Environment string
	Version     string
}

// getConfig returns application configuration from environment variables
func getConfig() *Config {
	return &Config{
		Port:        getEnv("PORT", defaultPort),
		Environment: getEnv("ENVIRONMENT", "development"),
		Version:     getEnv("VERSION", "1.0.0"),
	}
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// healthCheckHandler returns a health check endpoint
func healthCheckHandler(logger *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		checks := make(map[string]string)
		
		// Check database connection
		if db.DB != nil {
			if sqlDB, err := db.DB.DB(); err == nil {
				if err := sqlDB.Ping(); err == nil {
					checks["database"] = "healthy"
				} else {
					checks["database"] = "unhealthy: " + err.Error()
				}
			} else {
				checks["database"] = "unhealthy: " + err.Error()
			}
		} else {
			checks["database"] = "unhealthy: not connected"
		}

		// Determine overall status
		status := "healthy"
		for _, check := range checks {
			if check != "healthy" {
				status = "unhealthy"
				break
			}
		}

		response := HealthResponse{
			Status:    status,
			Timestamp: time.Now(),
			Version:   getEnv("VERSION", "1.0.0"),
			Checks:    checks,
		}

		w.Header().Set("Content-Type", "application/json")
		if status == "unhealthy" {
			w.WriteHeader(http.StatusServiceUnavailable)
		}

		if err := json.NewEncoder(w).Encode(response); err != nil {
			logger.Printf("Error encoding health check response: %v", err)
		}
	}
}

// Track server start time for uptime calculation
var serverStartTime = time.Now()

// metricsHandler returns basic metrics endpoint
func metricsHandler(logger *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Basic metrics - in production, you'd use Prometheus or similar
		metrics := map[string]interface{}{
			"uptime":    time.Since(serverStartTime).String(),
			"timestamp": time.Now(),
			"version":   getEnv("VERSION", "1.0.0"),
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(metrics); err != nil {
			logger.Printf("Error encoding metrics response: %v", err)
		}
	}
}

// loggingMiddleware logs HTTP requests
func loggingMiddleware(next http.Handler, logger *log.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		
		// Create a response writer wrapper to capture status code
		wrapper := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}
		
		next.ServeHTTP(wrapper, r)
		
		duration := time.Since(start)
		logger.Printf("%s %s %d %s %s", 
			r.Method, 
			r.URL.Path, 
			wrapper.statusCode, 
			duration, 
			r.RemoteAddr,
		)
	})
}

// responseWriter wraps http.ResponseWriter to capture status code
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

// recoveryMiddleware recovers from panics
func recoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic recovered: %v", err)
				http.Error(w, "Internal server error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

// corsMiddleware adds CORS headers to the response
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// In production, you should specify allowed origins
		origin := r.Header.Get("Origin")
		if origin == "" {
			origin = "*"
		}
		
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, X-Requested-With")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Max-Age", "86400")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// Add security headers
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")

		next.ServeHTTP(w, r)
	})
}

// graphqlMiddleware adds GraphQL-specific middleware
func graphqlMiddleware(logger *log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)
			duration := time.Since(start)
			
			if duration > 5*time.Second {
				logger.Printf("Slow GraphQL query detected: %s took %s", r.URL.Path, duration)
			}
		})
	}
} 