package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/Imjaopedro/firstGoProject/models"
)

type TaskHandler struct {
	DB *sql.DB
}

// Constructr for TaskHandler
func NewTaskHandler(db *sql.DB) *TaskHandler {
	return &TaskHandler{DB: db}
	// return new TaskHandler(db)   -em Java
	// DB: db = this.DB = db
}

func (TaskHandler *TaskHandler) ReadTasks(writer http.ResponseWriter, request *http.Request) {

	rows, err := TaskHandler.DB.Query("SELECT * FROM tasks")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		tasks = append(tasks, task)

	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(tasks)

}
