package main

import (
	"github.com/gin-gonic/gin"
	"github.com/viviyee/go-crud/controllers"
	"github.com/viviyee/go-crud/initializers"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/posts", controllers.GetPosts)
	r.POST("/posts", controllers.StorePost)
	r.GET("/posts/:id", controllers.GetPost)
	r.PUT("/posts/:id", controllers.UpdatePost)
	r.DELETE("/posts/:id", controllers.DeletePost)

	r.Run() // listen and serve on 0.0.0.0:8080
}
