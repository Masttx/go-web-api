package main

import (
	"fmt"
	"projetoinfiel/internal/database"
	"projetoinfiel/internal/repositories"
)

func main() {
	fmt.Println("Iniciando servidor")

	db := database.NewMySQLConnection()
	fmt.Println("Banco conectado")

	userRepository := repositories.NewUserRepository(db)

	//_, err := userRepository.Create("Mateus", "askdasdk@gmsadasdl")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	users, err := userRepository.List()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(users)
}
