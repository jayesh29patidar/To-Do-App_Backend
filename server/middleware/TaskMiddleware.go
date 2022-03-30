package middleware

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"

// 	"github.com/gorilla/mux"
// 	"github.com/jayesh29patidar/golang-react-todo/database"
// 	"github.com/jayesh29patidar/golang-react-todo/models"
// )

// func GetTaskByUser(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Access-Control-Allow-Methods", "POST")
// 	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

// 	vars := mux.Vars(r)
// 	userID, err := vars["userID"]
// 	if !err {
// 		fmt.Println("id is missing in parameters")
// 	}

// 	fmt.Println("Geting task with userID", userID)
// 	var tasksList = database.getTasksByID(userID)

// 	if tasksList == nil {
// 		tasksList = make([]models.ToDoList, 0)
// 	}
// 	json.NewEncoder(w).Encode(tasksList)
// }

// func GetAllTask(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/x-www-form urlencoded")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	payload := database.getAllTasks()
// 	json.NewEncoder(w).Encode(payload)
// }

// func CreateTask(w http.ResponseWriter, r *http.Request) {

// 	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Access-Control-Allow-Methods", "POST")
// 	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
// 	var task models.ToDoList

// 	_ = json.NewDecoder(r.Body).Decode(&task)

// 	var taskCreated = database.insertOneTask(task)
// 	json.NewEncoder(w).Encode(taskCreated)
// 	// database.insertOneTask(task)
// 	// json.NewEncoder(w).Encode(task)

// }

// func TaskComplete(w http.ResponseWriter, r *http.Request) {

// 	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Access-Control-Allow-Methods", "PUT")
// 	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

// 	params := mux.Vars(r)
// 	database.taskComplete(params["id"])
// 	json.NewEncoder(w).Encode(params["id"])
// }

// func UndoTask(w http.ResponseWriter, r *http.Request) {

// 	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Access-Control-Allow-Methods", "PUT")
// 	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

// 	params := mux.Vars(r)
// 	database.undoTask(params["id"])
// 	json.NewEncoder(w).Encode(params["id"])
// }

// func DeleteTask(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
// 	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
// 	params := mux.Vars(r)
// 	database.deleteOneTask(params["id"])
// 	json.NewEncoder(w).Encode(params["id"])
// }

// func DeleteAllTask(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	count := database.deleteAllTask()
// 	json.NewEncoder(w).Encode(count)

// }
