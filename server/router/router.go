package router

import (
	"github.com/gorilla/mux"
	"github.com/jayesh29patidar/golang-react-todo/middleware"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/user/add", middleware.AddUser).Methods("POST", "OPTIONS")
	router.HandleFunc("/user/login", middleware.LoginUser).Methods("POST", "OPTIONS")

	router.HandleFunc("/api/task", middleware.GetAllTask).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/task", middleware.CreateTask).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/tasksByUser/{userID}", middleware.GetTaskByUser).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/task/{id}", middleware.TaskComplete).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/undoTask/{id}", middleware.UndoTask).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/deleteTask/{id}", middleware.DeleteTask).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/deleteTask/{id}", middleware.DeleteTask).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/updateTask/{id}", middleware.UpdateTask).Methods("PUT", "OPTIONS")
	return router
}
