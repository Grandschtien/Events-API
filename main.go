package main

import (
	"crypto/tls"
	authHandlers "events-api/authentication/handlers"
	"events-api/authentication/middleware"
	"events-api/core/db"
	eventHandlers "events-api/events/handlers"
	"fmt"
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

	err, eventDB := db.SetupEventDB()
	err, usersDB := db.SetupUserDB()

	if err != nil {
		fmt.Println(err.Error())
		panic("DB did not initialize")
	}

	certPath := os.Getenv("CERTPATH")
	keyPath := os.Getenv("KEYPATH")
	cert, err := tls.LoadX509KeyPair(certPath, keyPath)

	if err != nil {
		log.Fatalf("Failed to load certificates: %v", err)
	}

	handlers := eventHandlers.EventHandlers{DB: eventDB}

	defer eventDB.Close()
	defer usersDB.Close()

	// Setup events handlers
	router.GET("/events", handlers.GetEvents)
	router.GET("/event", handlers.GetEvent)
	router.Use(middleware.AuthMiddleware()).DELETE("/event", handlers.DeleteEvent)
	router.Use(middleware.AuthMiddleware()).POST("/event", handlers.AddEvent)

	// Setup auth handlers
	authGroup := router.Group("/auth")

	authHandlers := authHandlers.AuthHandlers{DB: usersDB}

	authGroup.GET("/login", authHandlers.LoginUser)
	authGroup.POST("/registration", authHandlers.RegisterUser)

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
