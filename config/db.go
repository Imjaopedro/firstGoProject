package config

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Imjaopedro/firstGoProject/models"
	_ "modernc.org/sqlite"
)

// retorno um ponteiro na memoria para o DB
// conceito de Singleton
func SetUpDatabase() *sql.DB {
	dbPath := "./meubanco.db"
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// reset da tabela no início, não no shutdown
	_, err = db.Exec("DROP TABLE IF EXISTS tasks")
	if err != nil {
		log.Println("Erro ao dropar tabela:", err)
	}

	_, err = db.Exec(models.CreateTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conexão com SQLite aberta e tabela tasks criada")
	return db
}
