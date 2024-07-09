package main

import (
	"go-api-server/controllers"
	"go-api-server/initializers"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}
func main() {
	r := gin.Default()
	r.POST("/posts", controllers.PostCreate)
	r.GET("/posts", controllers.GetPosts)
	r.GET("/posts/:id", controllers.GetPost)
	r.PUT("/posts/", controllers.UpdatePost)
	r.DELETE("/posts/:id", controllers.DeletePost)
	port := os.Getenv("PORT")
	r.Run(port)
}
