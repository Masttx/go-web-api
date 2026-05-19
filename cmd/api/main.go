package main

import (
	"fmt"
	"log"
	"net/http"
	"projetoinfiel/internal/api"
	"projetoinfiel/internal/database"
	"projetoinfiel/internal/repositories"
)

func main() {
	fmt.Println("Iniciando servidor")

	db := database.NewMySQLConnection()
	fmt.Println("Banco conectado")

	userRepository := repositories.NewUserRepository(db)

	userAPI := api.NewUserAPI(*userRepository)

	http.HandleFunc("POST /user", userAPI.Create)
	http.HandleFunc("PUT /user/{id}", userAPI.Update)
	http.HandleFunc("GET /user/{id}", userAPI.Read)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}
