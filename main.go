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

	err, postgres := db.SetupEventDB()

	if err != nil {
		panic("DB did not initialize")
	}

	handlers := handlers.Handlers{DB: postgres}

	defer postgres.Close()

	router.GET("/events", handlers.GetEvents)
	router.GET("/event", handlers.GetEvent)
	router.DELETE("/event", handlers.DeleteEvent)
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
