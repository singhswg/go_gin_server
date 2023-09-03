package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/singhswg/go_gin_server/database"
)

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ping received",
	})
}

func getUsers(c *gin.Context) { // Read the database and get all users
	fmt.Println("Reading to DB now...")
	database.Connect("read", c)
}

func getUserInfo(c *gin.Context) { // Query the database and get user info
	fmt.Println("getting info of user", c.Param("name"))
	database.Connect("getUserInfo", c)
}

func addUser(c *gin.Context) { // Handled POST request that accepts JSON data
	fmt.Println("Adding user...")
	database.Connect("addUser", c)
}
