package database

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"

	"github.com/jayesh29patidar/golang-react-todo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func InsertUser(user models.User) models.User {
	EncryptedPassword, err := HashPassword(user.Password)
	user.Password = EncryptedPassword
	if err != nil {
		log.Fatal(err)
	}
	insertResult, err := userCollection.InsertOne(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Insterted User", insertResult.InsertedID, user.Username)
	var result models.User
	var nullUser models.User
	err = userCollection.FindOne(context.Background(), bson.M{"_id": insertResult.InsertedID}).Decode(&result)
	fmt.Println("RESULT ", result.ID, result.Username)
	fmt.Println("RESULT ", result)
	if err != nil {
		return nullUser
	}
	fmt.Println("Done Succ No error ")
	return result
}

func TokenGenerator() string {
	b := make([]byte, 8)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func LoginCheck(user models.User) models.User {
	var result models.User
	var nullUser models.User
	err := userCollection.FindOne(context.Background(), bson.M{"username": user.Username}).Decode(&result)
	if err != nil {
		return nullUser
	}
	var passwordMatch = CheckPasswordHash(user.Password, result.Password)
	if !passwordMatch {
		return nullUser
	}
	fmt.Println("Matched Data", result.ID, result.Username, result.Password)
	return result
}

func GetTasksByID(userIDHex string) []models.ToDoList {
	userID, err := primitive.ObjectIDFromHex(userIDHex)
	if err != nil {
		panic(err)
	}
	var tasksList []models.ToDoList
	findResult, err := taskCollection.Find(context.TODO(), bson.M{"user._id": userID})
	if err != nil {
		log.Fatal(err)
		return tasksList
	}

	for findResult.Next(context.TODO()) {
		var task models.ToDoList
		err := findResult.Decode(&task)
		if err != nil {
			log.Fatal(err)
		}
		tasksList = append(tasksList, task)
	}
	if err := findResult.Err(); err != nil {
		log.Fatal(err)
	}
	return tasksList
}

func GetAllTasks() []primitive.M {
	cur, err := taskCollection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var results []primitive.M
	for cur.Next(context.Background()) {
		var result bson.M
		e := cur.Decode(&result)
		if e != nil {
			log.Fatal(e)
		}
		results = append(results, result)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(context.Background())
	return results
}

func InsertOneTask(task models.ToDoList) {
	insertResult, err := taskCollection.InsertOne(context.Background(), task)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Insterted a single record", insertResult.InsertedID)
}

func CallTaskComplete(task string) {
	fmt.Println(task)
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": true}}
	result, err := taskCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("modified count: ", result.ModifiedCount)
}

func CallUndoTask(task string) {
	fmt.Println(task)
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": false}}
	result, err := taskCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("modified count: ", result.ModifiedCount)
}

func DeleteOneTask(task string) {
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	d, err := taskCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted Document: ", d.DeletedCount)
}

func UpdateTaskDatabase(taskIdHex string, updatedTask models.ToDoList) {
	taskId, err := primitive.ObjectIDFromHex(taskIdHex)
	if err != nil {
		panic(err)
	}
	fmt.Println("taskId", taskId)
	filter := bson.M{"_id": taskId}
	result, err := taskCollection.ReplaceOne(context.TODO(), filter, updatedTask)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(
		"insert: %d, updated%d, deleted: %d /n",
		result.MatchedCount,
		result.ModifiedCount,
		result.UpsertedCount,
	)
}
