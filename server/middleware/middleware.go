package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jayesh29patidar/golang-react-todo/models"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"crypto/rand"

	"golang.org/x/crypto/bcrypt"
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

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func AddUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	fmt.Println(r.Method)
	if r.Method == "OPTIONS" {
		return
	}

	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	fmt.Println("user done", user)
	if user.Username == "" {
		return
	}
	var listedUser = InsertUser(user)
	json.NewEncoder(w).Encode(listedUser)
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

func LoginUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	token := TokenGenerator()
	fmt.Println("login Token", token)

	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	fmt.Println("login hit with user", user)

	var loggedUser = LoginCheck(user)
	var userWithToken models.UserWithToken
	userWithToken.Token = token
	userWithToken.User = loggedUser

	json.NewEncoder(w).Encode(userWithToken)
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

func GetTaskByUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	vars := mux.Vars(r)
	userID, err := vars["userID"]
	if !err {
		fmt.Println("id is missing in parameters")
	}

	fmt.Println("Geting task with userID", userID)
	var tasksList = getTasksByID(userID)

	if tasksList == nil {
		tasksList = make([]models.ToDoList, 0)
	}
	json.NewEncoder(w).Encode(tasksList)
}

func getTasksByID(userIDHex string) []models.ToDoList {
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

func GetAllTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	payload := getAllTasks()
	json.NewEncoder(w).Encode(payload)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == "OPTIONS" {
		return
	}
	var task models.ToDoList

	_ = json.NewDecoder(r.Body).Decode(&task)
	fmt.Println("Task create hit ", task)

	insertOneTask(task)
	json.NewEncoder(w).Encode(task)

}

func TaskComplete(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	taskComplete(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func UndoTask(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	undoTask(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	params := mux.Vars(r)
	deleteOneTask(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func getAllTasks() []primitive.M {
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

func insertOneTask(task models.ToDoList) {
	insertResult, err := taskCollection.InsertOne(context.Background(), task)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Insterted a single record", insertResult.InsertedID)
}

func taskComplete(task string) {
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

func undoTask(task string) {
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

func deleteOneTask(task string) {
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	d, err := taskCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted Document: ", d.DeletedCount)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		return
	}
	vars := mux.Vars(r)
	id, err := vars["id"]
	if !err {
		fmt.Println("id missing")
	}
	fmt.Println(`update id := `, id)
	var task models.ToDoList
	fmt.Println("update task hit with task ", task)
	json.NewDecoder(r.Body).Decode(&task)
	UpdateTaskDatabase(id, task)

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
