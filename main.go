package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/lordofthemind/jsonCrudGorm/controllers"
	"github.com/lordofthemind/jsonCrudGorm/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToPostgreSQL()
}

func main() {
	r := gin.Default()

	api := r.Group("/api/v1")
	{
		api.POST("/posts", controllers.CreatePost)
		api.GET("/posts", controllers.GetPosts)
		api.GET("/posts/:id", controllers.GetPost)
		api.PUT("/posts/:id", controllers.UpdatePost)
		api.DELETE("/posts/:id", controllers.DeletePost)
	}

	r.Run("localhost:" + os.Getenv("PORT"))
}
