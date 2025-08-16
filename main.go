package main

import (
	"log"
	"net/http"

	"github.com/Imjaopedro/firstGoProject/config"
	"github.com/Imjaopedro/firstGoProject/handlers"
	"github.com/Imjaopedro/firstGoProject/models"
	"github.com/gorilla/mux"
)

// Function called when program start
// initialize connection with db
// map the endpoints
func main() {

	dbConnection := config.SetUpDatabase()
	defer dbConnection.Close()

	_, err := dbConnection.Exec(models.CreateTableSQL)

	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()
	TaskHandler := handlers.NewTaskHandler(dbConnection)

	router.HandleFunc("/tasks", TaskHandler.ReadTasks).Methods("GET")
	router.HandleFunc("/tasks", TaskHandler.CreateTask).Methods("POST")
	router.HandleFunc("/task/{id}", TaskHandler.UpdateTask).Methods("PUT")
	router.HandleFunc("/task/{id}", TaskHandler.DeleteTask).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3030", router))

}
