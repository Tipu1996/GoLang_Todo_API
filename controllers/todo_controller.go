package controllers

import (
	"example/libraryAPI/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddTodo(context *gin.Context, client *mongo.Client) {
	var newTodo models.Todo
	newTodo.ID = primitive.NewObjectID()

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	coll := client.Database("GoLang").Collection("Todo")

	// Insert the new todo into the collection.
	result, err := coll.InsertOne(context, newTodo)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add the todo"})
		return
	}

	// Get the ID of the newly created todo.
	newTodoID := result.InsertedID

	// Respond to the user with the ID of the newly created todo.
	context.JSON(http.StatusOK, gin.H{"message": "Todo added", "id": newTodoID})
}

func GetTodos(context *gin.Context, client *mongo.Client) {
	coll := client.Database("GoLang").Collection("Todo")

	filter := bson.D{{}}

	cur, err := coll.Find(context, filter)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add the todo"})
		return
	}

	var todos []models.Todo

	for cur.Next(context) {
		var todo models.Todo

		// Decode the other fields from the MongoDB document.
		if err := cur.Decode(&todo); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode todos"})
			return
		}

		todos = append(todos, todo)
	}

	cur.Close(context)

	context.IndentedJSON(http.StatusOK, todos)
}

func GetTodoById(context *gin.Context, client *mongo.Client) {
	idParam := context.Param("id")
	objID, err := primitive.ObjectIDFromHex(idParam)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	coll := client.Database("GoLang").Collection("Todo")

	filter := bson.D{{"_id", objID}}

	var todo models.Todo

	err = coll.FindOne(context, filter).Decode(&todo)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)
}

func UpdateTodoById(context *gin.Context, client *mongo.Client) {
	idParam := context.Param("id")
	objID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	coll := client.Database("GoLang").Collection("Todo")
	filter := bson.D{{"$set", bson.D{{"completed", true}}}}

	result, err := coll.UpdateByID(context, objID, filter)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}
	if result.ModifiedCount == 0 {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update the todo"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Todo updated successfully"})
}

func DelTodo(context *gin.Context, client *mongo.Client) {
	idParam := context.Param("id")
	objID, err := primitive.ObjectIDFromHex(idParam)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	coll := client.Database("GoLang").Collection("Todo")

	filter := bson.D{{"_id", objID}}

	var todo models.Todo

	err = coll.FindOneAndDelete(context, filter).Decode(&todo)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)
}
