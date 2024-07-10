package main

import (
	"go-api-server/controllers"
	"go-api-server/initializers"
	"go-api-server/middlewares"
	"go-api-server/migrate"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	migrate.Migrate()
}
func main() {
	r := gin.Default()
	r.POST("/posts", controllers.PostCreate)
	r.GET("/posts", controllers.GetPosts)
	r.GET("/posts/:id", controllers.GetPost)
	r.PUT("/posts/", controllers.UpdatePost)
	r.DELETE("/posts/:id", controllers.DeletePost)

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middlewares.RequireAuth, middlewares.RoleAuth, controllers.Validate)
	port := os.Getenv("PORT")
	r.Run(port)
}
