package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jayesh29patidar/golang-react-todo/database"
	"github.com/jayesh29patidar/golang-react-todo/models"

	"github.com/gorilla/mux"
)

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
	var tasksList = database.GetTasksByID(userID)

	if tasksList == nil {
		tasksList = make([]models.ToDoList, 0)
	}
	json.NewEncoder(w).Encode(tasksList)
}

func GetAllTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	payload := database.GetAllTasks()
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

	database.InsertOneTask(task)
	json.NewEncoder(w).Encode(task)

}

func TaskComplete(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	database.CallTaskComplete(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func UndoTask(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	database.CallUndoTask(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	params := mux.Vars(r)
	database.DeleteOneTask(params["id"])
	json.NewEncoder(w).Encode(params["id"])
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
	database.UpdateTaskDatabase(id, task)
}
