package main

import (
	"log"
	"net/http"

	"github.com/Imjaopedro/firstGoProject/config"
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
	router.HandleFunc("/tasks", func(w http.ResponseWriter,
		r *http.Request) {

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!"))
	}).Methods("GET")

	log.Fatal(http.ListenAndServe(":3030", router))

}
