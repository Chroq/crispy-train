package main

import (
	"log"
	"net/http"

	"backathon/internal/interfaces/rest"
	"backathon/pkg"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"state": "up",
		})
	})

	// copy swagger.yml to static/swaggerui
	err := pkg.CopyFile("./api/swagger.yml", "./static/swaggerui/swagger.yml")
	if err != nil {
		log.Fatalf("Failed to copy swagger.yml: %v", err)
	}
	r.Static("/swagger", "./static/swaggerui")

	petStoreHandler := rest.NewPetStoreHandler()

	server := rest.NewServerHandler(petStoreHandler)

	rest.RegisterHandlers(r, server)

	log.Fatal(r.Run())
}
