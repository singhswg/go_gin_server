package database

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	db_name = "postgres"
	host    = "localhost"
	port    = 5432
)

type user struct {
	Name  string `json:"name"`
	Title string `json:"title"`
	City  string `json:"city"`
}

func Connect(operation string, c *gin.Context) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, os.Getenv("PG_USER"), os.Getenv("PG_PASS"), db_name)
	fmt.Printf(connStr + "\n")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")

	if operation == "read" {
		ReadDB(db, c)
	} else if operation == "getUserInfo" {
		getUserInfo(db, c)
	} else if operation == "addUser" {
		addUser(db, c)
	} else {
		c.IndentedJSON(http.StatusOK, "Incorrect Operation selected")
	}
}

func addUser(db *sql.DB, c *gin.Context) {

	var addUserData map[string]interface{} // Parse the JSON data from request body
	if err := c.BindJSON(&addUserData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for key, value := range addUserData {
		println(key, value)
	}

	fmt.Println("Adding new user", addUserData["name"])
	fmt.Println("Adding new user's title", addUserData["title"])

	_, err := db.Exec("INSERT INTO users (name, title, city) VALUES ($1, $2, $3)", addUserData["name"], addUserData["title"], addUserData["city"]) // Insert to table                                                                                                                                             // Run query
	if err != nil {
		panic(err)
	}
	fmt.Println("InsertDB executed")

	c.JSON(http.StatusCreated, gin.H{"data": "User was added"})
}

func ReadDB(db *sql.DB, c *gin.Context) {
	rows, err := db.Query("SELECT * from users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var col1, col2, col3 string
		err := rows.Scan(&col1, &col2, &col3)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			panic(err)
		}

		data := map[string]interface{}{
			"Name":  col1,
			"Title": col2,
			"City":  col3,
		}
		c.IndentedJSON(http.StatusOK, data)
	}
	fmt.Println("ReadDB executed")
}

func getUserInfo(db *sql.DB, c *gin.Context) { // Function to execute query on user
	name := c.Param("name")
	rows, err := db.Query("SELECT * from users where name = $1", name)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var col1, col2, col3 string
		err := rows.Scan(&col1, &col2, &col3)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			panic(err)
		}
		c.IndentedJSON(http.StatusOK, "User found")
	} else {
		c.IndentedJSON(http.StatusOK, "User not found")
	}
}
