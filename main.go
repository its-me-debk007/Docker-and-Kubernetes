package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func main() {
	//database.ConnectDatabase()

	app := gin.Default()
	app.Use(gin.Recovery())

	app.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})

	port := os.Getenv("PORT")

	if err := app.Run(":" + port); err != nil {
		log.Fatal("App listen error:-\n" + err.Error())
	}

	log.Println("Server running at port: ", port)
}
