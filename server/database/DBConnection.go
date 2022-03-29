// package database

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"os"
// 	"github.com/joho/godotenv"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// func init() {
// 	loadTheEnv()
// 	createDBInstance()
// }

// func loadTheEnv() {
// 	err := godotenv.Load(".env")
// 	if err != nil {
// 		log.Fatal("Error loading the .env file")
// 	}
// }

// func createDBInstance() {
// 	connectionString := os.Getenv("DB_URI")
// 	dbName := os.Getenv("DB_NAME")
// 	collName := os.Getenv("DB_COLLECTION_NAME")

// 	clientOptions := options.Client().ApplyURI(connectionString)

// 	client, err := mongo.Connect(context.TODO(), clientOptions)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	err = client.Ping(context.TODO(), nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("connect to mongodb!")

// 	collection = client.Database(dbName).Collection(collName)
// 	fmt.Println("collection instance created")
// }
