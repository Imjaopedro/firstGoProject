package config

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

// retorno um ponteiro na memoria para o DB
// conceito de Singleton
func SetUpDatabase() *sql.DB {
	dbPath := "./meubanco.db"
	connString := fmt.Sprintf("%s", dbPath)

	// driver correto para modernc.org/sqlite
	db, err := sql.Open("sqlite", connString)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Conexão com SQLite aberta com Sucesso")
	return db
}
