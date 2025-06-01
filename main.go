package main

import (
	"events-api/db"
	"events-api/handlers"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	postgress := db.Setup("events")

	handlers := handlers.Handlers{DB: &db.DB{DB: postgress}}

	router.GET("/events", handlers.GetEvents)
	router.GET("/event/:id", handlers.GetEvent)
	router.DELETE("/events/:id", handlers.DeleteEvent)
	router.POST("/event", handlers.AddEvent)

	server := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}
