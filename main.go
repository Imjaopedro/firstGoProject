package main

import (
	"log"
	"net/http"

	"github.com/Imjaopedro/firstGoProject/config"
)

// Function called when program start
// initialize connection with db
// map the endpoints
func main() {

	dbConnection := config.SetUpDatabase()

	defer dbConnection.Close()

	log.Fatal(http.ListenAndServe(":3030", nil))

}
