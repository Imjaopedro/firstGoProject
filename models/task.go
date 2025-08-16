package models

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:description`
	Status      bool   `json:"status"`
}

const (
	TableName      = "tasks"
	CreateTableSQL = `
	CREATE TABLE IF NOT EXISTS tasks(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    description TEXT,
    status BOOLEAN NOT NULL DEFAULT false
);

	`
)
