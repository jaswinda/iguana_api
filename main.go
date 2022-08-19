package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		c.String(http.StatusOK, "Hello %s", firstname)
	})
	router.GET("/jaswindar", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		c.String(http.StatusOK, "Hello from jaswindar %s", firstname)
	})
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello to the ")
	})

	router.Run(":8000")
}
