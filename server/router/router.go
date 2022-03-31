package router

import (
	"github.com/gorilla/mux"
	"github.com/jayesh29patidar/golang-react-todo/service"
)

func Router() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/user/add", service.AddUser).Methods("POST", "OPTIONS")
	router.HandleFunc("/user/login", service.LoginUser).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/task", service.GetAllTask).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/task", service.CreateTask).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/tasksByUser/{userID}", service.GetTaskByUser).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/task/{id}", service.TaskComplete).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/undoTask/{id}", service.UndoTask).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/deleteTask/{id}", service.DeleteTask).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/deleteTask/{id}", service.DeleteTask).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/updateTask/{id}", service.UpdateTask).Methods("PUT", "OPTIONS")

	return router
}
