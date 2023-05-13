package main

import (
	"github.com/gin-gonic/gin"
	"github.com/its-me-debk007/Docker-and-Kubernetes/database"
	"github.com/its-me-debk007/Docker-and-Kubernetes/model"
	"log"
	"net/http"
	"os"
)

func main() {
	database.ConnectDatabase()

	app := gin.Default()
	app.Use(gin.Recovery())

	app.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})

	app.POST("/user", func(c *gin.Context) {
		user := model.User{
			"Ajay",
			"Kumar",
			"kumarajay@gmail.com",
			23,
		}

		database.DB.Save(&user)

		c.JSON(http.StatusAccepted, gin.H{
			"message": "added user successfully",
		})
	})

	port := os.Getenv("PORT")

	if err := app.Run(":" + port); err != nil {
		log.Fatal("App listen error:-\n" + err.Error())
	}

	log.Println("Server running at port: ", port)
}
