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

	areaRepository := repositories.NewAreaRepository(db)
	userRepository := repositories.NewUserRepository(db)

	areaAPI := api.NewAreaAPI(*areaRepository)
	userAPI := api.NewUserAPI(*userRepository)

	http.HandleFunc("POST /area", areaAPI.Create)
	http.HandleFunc("PATCH /area/{id}", areaAPI.Update)
	http.HandleFunc("GET /area/{id}", areaAPI.Read)

	http.HandleFunc("POST /user", userAPI.Create)
	http.HandleFunc("PATCH /user/{id}", userAPI.Update)
	http.HandleFunc("GET /user/{id}", userAPI.Read)
	http.HandleFunc("GET /user/area/{area_id}", userAPI.ListByArea)
	http.HandleFunc("PATCH /user/area/{id}", userAPI.UpdateUserArea)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}
