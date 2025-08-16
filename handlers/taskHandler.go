package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/Imjaopedro/firstGoProject/models"
	"github.com/gorilla/mux"
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
func (taskHandler *TaskHandler) CreateTask(writer http.ResponseWriter, request *http.Request) {
	var task models.Task

	// Decodifica o JSON do corpo da requisição
	err := json.NewDecoder(request.Body).Decode(&task)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	// Insere no banco
	_, err = taskHandler.DB.Exec(
		`INSERT INTO tasks (title, description, status) VALUES ($1, $2, $3)`,
		task.Title, task.Description, task.Status,
	)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(map[string]string{
		"message": "Task criada com sucesso!",
	})
}

func (taskHandler *TaskHandler) UpdateTask(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	var task models.Task
	err := json.NewDecoder(request.Body).Decode(&task)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = taskHandler.DB.Exec(
		`UPDATE tasks SET title=$1, description=$2, status=$3 WHERE id=$4`,
		task.Title, task.Description, task.Status, id,
	)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(map[string]string{
		"message": "Task atualizada com sucesso!",
	})
}

func (taskHandler *TaskHandler) DeleteTask(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	_, err := taskHandler.DB.Exec(
		`DELETE FROM tasks WHERE id=$1`, id,
	)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(map[string]string{
		"message": "Task deletada com sucesso!",
	})
}
