package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.Default()
	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		c.String(http.StatusOK, "Hello from jaswindar %s", firstname)
	})
	router.GET("/jaswindar", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		c.String(http.StatusOK, "Hello from jaswindar %s", firstname)
	})
	router.GET("/", func(c *gin.Context) {
		//ready .env file
		Host := goDotEnvVariable("DB_USER")

		c.String(http.StatusOK, "Hello from jaswindar %s", Host)
	})

	router.Run(":9000")
}
func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
