package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

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

func handleShutdown(db *sql.DB) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c // espera o sinal

	fmt.Println("\nShutting down... Dropping table tasks")
	_, err := db.Exec("DROP TABLE IF EXISTS tasks")
	if err != nil {
		log.Println("Erro ao dropar a tabela tasks:", err)
	} else {
		log.Println("Tabela tasks deletada com sucesso!")
	}

	db.Close()
	os.Exit(0)
}
