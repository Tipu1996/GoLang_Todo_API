package main

import (
	"context"
	"example/libraryAPI/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// func connectToMongoDB() *mongo.Client {
// 	if err := godotenv.Load(); err != nil {
// 		log.Println("No .env file found")
// 	}
// 	uri := os.Getenv("MONGODB_URI")
// 	if uri == "" {
// 		log.Fatal("You must set your 'MONGODB_URI' environment variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
// 	}

// 	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Ping the MongoDB server to confirm the connection
// 	err = client.Ping(context.TODO(), nil)
// 	if err != nil {
// 		log.Fatal("Failed to ping MongoDB server")
// 	}

// 	// Log a message confirming the successful connection
// 	log.Println("Successfully connected to MongoDB server!")

// 	defer func() {
// 		if err := client.Disconnect(context.TODO()); err != nil {
// 			panic(err)
// 		}
// 	}()

// 	return client
// }

func main() {
	// client := connectToMongoDB()
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	// Ping the MongoDB server to confirm the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Failed to ping MongoDB server")
	}

	// Log a message confirming the successful connection
	log.Println("Successfully connected to MongoDB server!")

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	router := gin.Default()
	routes.SetupRoutes(router, client)
	router.Run("localhost:9090")
}
