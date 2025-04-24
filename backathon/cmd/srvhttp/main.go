package main

import (
	"log"
	"net/http"

	"backathon/internal/interfaces/rest"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"state": "up",
		})
	})

	// Initialize your handlers here
	petStoreHandler := rest.NewPetStoreHandler()
	server := rest.NewServerHandler(petStoreHandler)

	rest.RegisterHandlers(r, server)

	log.Fatal(r.Run())
}
