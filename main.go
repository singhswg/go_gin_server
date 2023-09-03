package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	Name  string `json:"name"`
	Title string `json:"title"`
	City  string `json:"city"`
}

var users = []user{
	{Name: "Balpreet", Title: "Tech Lead", City: "Chicago"},
}

func main() {
	router := gin.Default()

	router.GET("/ping", ping)

	router.GET("/users", getUsers)

	router.POST("/adduser", addUser)

	router.GET("/users/:name", getUserInfo)

	router.NoRoute(handleUnsupported)

	log.Printf("Starting server")
	router.Run() // Runs default on localhost:8080
}

func handleUnsupported(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"message": "Endpoint not found",
	})
}
