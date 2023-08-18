package main

import (
	"encoding/json"
	// "go/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

type event struct {
	ID          string `json: id`
	Title       string `json: title`
	Description string `json: description`
	// TODO: Date and time, img links etc
}

var firstPage = event{
	ID: "1", Title: "My First Page", Description: "Hello World!",
}

func main() {
	router := gin.Default()

	router.GET("/", getHomepage)

	router.Run("127.0.0.1:8080")
}

func getHomepage(c *gin.Context) {
	print("Got Request")
	print(json.Marshal(firstPage))
	// jdata, _ := json.Marshal(firstPage)
	c.JSON(http.StatusOK, "{}")
}
