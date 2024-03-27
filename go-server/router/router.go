package router

import {
	"go-server/middlewear"

	"github.com/gorilla/mux"
}

func Router () *mux.Router {
	
	router := mux.NewRouter()

	router.HandleFunc("/api/task", middlewear.GetAllTask).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/task", middlewear.CreateTask).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/task/{id}", middlewear.TaskComplete).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/undoTask/{id}", middlewear.UndoTask).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/deleteTask/{id}", middlewear.DeleteTask).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/api/deleteAllTask", middlewear.DeleteAllTask).Methods("DELETE", "OPTIONS")

	return router
}