package database

// import (
// 	"context"
// 	"crypto/rand"
// 	"fmt"
// 	"log"
// 	"os"

// 	"github.com/jayesh29patidar/golang-react-todo/models"
// 	"github.com/joho/godotenv"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// 	"golang.org/x/crypto/bcrypt"
// )

// var collection *mongo.Collection

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

// func HashPassword(password string) (string, error) {
// 	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
// 	return string(bytes), err
// }

// func CheckPasswordHash(password, hash string) bool {
// 	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
// 	return err == nil
// }

// func InsertUser(user models.User) models.User {
// 	EncryptedPassword, err := HashPassword(user.Password)
// 	user.Password = EncryptedPassword
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	insertResult, err := collection.InsertOne(context.Background(), user)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Insterted User", insertResult.InsertedID, user.Username)

// 	var result models.User
// 	var nullUser models.User
// 	err = collection.FindOne(context.Background(), bson.M{"_id": insertResult.InsertedID}).Decode(&result)
// 	fmt.Println("RESULT ", result.ID, result.Username)
// 	fmt.Println("RESULT ", result)
// 	if err != nil {
// 		return nullUser
// 	}
// 	fmt.Println("No ERR ")
// 	return result
// }

// func TokenGenerator() string {
// 	b := make([]byte, 8)
// 	rand.Read(b)
// 	return fmt.Sprintf("%x", b)
// }

// func LoginCheck(user models.User) models.User {

// 	var result models.User
// 	var nullUser models.User
// 	err := collection.FindOne(context.Background(), bson.M{"username": user.Username}).Decode(&result)
// 	if err != nil {
// 		return nullUser
// 	}

// 	var passwordMatch = CheckPasswordHash(user.Password, result.Password)
// 	if !passwordMatch {
// 		return nullUser
// 	}

// 	fmt.Println("Matched Data", result.ID, result.Username, result.Password)
// 	return result

// }

// func getTasksByID(userIDHex string) []models.ToDoList {
// 	userID, err := primitive.ObjectIDFromHex(userIDHex)
// 	if err != nil {
// 		panic(err)
// 	}
// 	var tasksList []models.ToDoList
// 	findResult, err := collection.Find(context.TODO(), bson.M{"user._id": userID})
// 	if err != nil {
// 		log.Fatal(err)
// 		return tasksList
// 	}

// 	for findResult.Next(context.TODO()) {
// 		var task models.ToDoList
// 		err := findResult.Decode(&task)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		tasksList = append(tasksList, task)
// 	}
// 	if err := findResult.Err(); err != nil {
// 		log.Fatal(err)
// 	}
// 	return tasksList
// }

// // ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// // ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// func getAllTasks() []primitive.M {
// 	cur, err := collection.Find(context.Background(), bson.D{{}})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	var results []primitive.M
// 	for cur.Next(context.Background()) {
// 		var result bson.M
// 		e := cur.Decode(&result)
// 		if e != nil {
// 			log.Fatal(e)
// 		}
// 		results = append(results, result)
// 	}
// 	if err := cur.Err(); err != nil {
// 		log.Fatal(err)
// 	}
// 	cur.Close(context.Background())
// 	return results
// }

// func insertOneTask(task models.ToDoList) {
// 	insertResult, err := collection.InsertOne(context.Background(), task)

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Insterted a single record", insertResult.InsertedID)
// }

// func taskComplete(task string) {
// 	fmt.Println(task)
// 	id, _ := primitive.ObjectIDFromHex(task)
// 	filter := bson.M{"_id": id}
// 	update := bson.M{"$set": bson.M{"status": true}}
// 	result, err := collection.UpdateOne(context.Background(), filter, update)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("modified count: ", result.ModifiedCount)
// }

// func undoTask(task string) {
// 	fmt.Println(task)
// 	id, _ := primitive.ObjectIDFromHex(task)
// 	filter := bson.M{"_id": id}
// 	update := bson.M{"$set": bson.M{"status": false}}
// 	result, err := collection.UpdateOne(context.Background(), filter, update)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("modified count: ", result.ModifiedCount)
// }

// func deleteOneTask(task string) {
// 	id, _ := primitive.ObjectIDFromHex(task)
// 	filter := bson.M{"_id": id}
// 	d, err := collection.DeleteOne(context.Background(), filter)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Deleted Document: ", d.DeletedCount)

// }

// func deleteAllTask() int64 {
// 	d, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Deleted Document", d.DeletedCount)
// 	return d.DeletedCount
// }

// func getUser() []primitive.M {
// 	cur, err := collection.Find(context.Background(), bson.D{{}})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	var results2 []primitive.M
// 	for cur.Next(context.Background()) {
// 		var result bson.M
// 		e := cur.Decode(&result)
// 		if e != nil {
// 			log.Fatal(e)
// 		}
// 		results2 = append(results2, result)
// 	}
// 	if err := cur.Err(); err != nil {
// 		log.Fatal(err)
// 	}
// 	cur.Close(context.Background())
// 	return results2
// }
