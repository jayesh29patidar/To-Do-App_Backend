package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var userCollection *mongo.Collection
var taskCollection *mongo.Collection

func init() {
	loadTheEnv()
	createDBInstance()
}

func loadTheEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading the .env file")
	}
}

func createDBInstance() {
	godotenv.Load()
	connectionString := os.Getenv("DB_URI")
	dbName := os.Getenv("DB_NAME")
	userCollName := os.Getenv("DB_COLLECTION_USER")
	taskCollName := os.Getenv("DB_COLLECTION_TASK")

	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connect to mongodb!")

	userCollection = client.Database(dbName).Collection(userCollName)
	taskCollection = client.Database(dbName).Collection(taskCollName)

	fmt.Println("collection instance created")
}
