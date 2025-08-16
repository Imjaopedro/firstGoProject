package main

import (
	"fmt"
	"net/http"

	"github.com/Imjaopedro/firstGoProject/cmd/database"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.NewConnection(":memory:")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repo := database.NewUserRepository(db)

	g := gin.Default()

	// Criar usuário
	g.POST("/users", func(ctx *gin.Context) {
		var user database.User
		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := repo.Create(&user); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, user)
	})

	// Listar usuários
	g.GET("/users", func(ctx *gin.Context) {
		users, err := repo.FindAll()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, users)
	})

	// Buscar usuário por ID
	g.GET("/users/:id", func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		var id int64
		fmt.Sscan(idParam, &id)

		user, err := repo.FindByID(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		ctx.JSON(http.StatusOK, user)
	})

	// Deletar usuário
	g.DELETE("/users/:id", func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		var id int64
		fmt.Sscan(idParam, &id)

		if err := repo.Delete(id); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.Status(http.StatusNoContent)
	})

	g.Run(":3000")
}
