package routes

import (
	"example/libraryAPI/controllers"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRoutes(router *gin.Engine, client *mongo.Client) {
	todoGroup := router.Group("/todos")
	{
		todoGroup.GET("/", func(c *gin.Context) {
			controllers.GetTodos(c, client) // Pass the client to the AddTodo function
		})
		todoGroup.GET("/:id", func(c *gin.Context) {
			controllers.GetTodoById(c, client) // Pass the client to the AddTodo function
		})
		todoGroup.PATCH("/:id", func(c *gin.Context) {
			controllers.UpdateTodoById(c, client) // Pass the client to the AddTodo function
		})
		todoGroup.DELETE("/:id", func(c *gin.Context) {
			controllers.DelTodo(c, client) // Pass the client to the AddTodo function
		})
		todoGroup.POST("/", func(c *gin.Context) {
			controllers.AddTodo(c, client) // Pass the client to the AddTodo function
		})
	}
}
