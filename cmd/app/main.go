package main

import (
	"net/http"

	"ruchka/internal/database"
	"ruchka/internal/handlers"
	"ruchka/internal/taskService"

	"github.com/gorilla/mux"
)

func main() {
	database.InitDB()
	database.DB.AutoMigrate(&taskService.Task{})

	repo := taskService.NewTaskRepository(database.DB)
	service := taskService.NewService(repo)

	handler := handlers.NewHandler(service)

	router := mux.NewRouter()
	router.HandleFunc("/api/task", handler.GetTasksHandler).Methods("GET")
	router.HandleFunc("/api/task", handler.PostTaskHandler).Methods("POST")
	router.HandleFunc("/api/task/{id}", handler.PatchTaskHandler).Methods("PATCH")
	router.HandleFunc("/api/task/{id}", handler.DeleteTaskHandler).Methods("DELETE")
	http.ListenAndServe(":8080", router)
}
