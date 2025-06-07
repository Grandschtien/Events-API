package main

import (
	"crypto/tls"
	"events-api/db"
	"events-api/handlers"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("config/.env")

	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	router := gin.Default()

	err, postgres := db.SetupEventDB()

	if err != nil {
		panic("DB did not initialize")
	}

	certPath := os.Getenv("CERTPATH")
	keyPath := os.Getenv("KEYPATH")
	cert, err := tls.LoadX509KeyPair(certPath, keyPath)

	if err != nil {
		log.Fatalf("Failed to load certificates: %v", err)
	}

	handlers := handlers.Handlers{DB: postgres}

	defer postgres.Close()

	router.GET("/events", handlers.GetEvents)
	router.GET("/event", handlers.GetEvent)
	router.DELETE("/event", handlers.DeleteEvent)
	router.POST("/event", handlers.AddEvent)

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	server := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		TLSConfig:      tlsConfig,
	}

	server.ListenAndServeTLS(certPath, keyPath)
}
