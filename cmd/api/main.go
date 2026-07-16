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
	areaUserRepository := repositories.NewAreaUserRepository(db)
	userRepository := repositories.NewUserRepository(db)

	areaAPI := api.NewAreaAPI(*areaRepository, *areaUserRepository)
	userAPI := api.NewUserAPI(*userRepository, *areaUserRepository)

	http.HandleFunc("POST /area", areaAPI.Create)
	http.HandleFunc("PATCH /area/{id}", areaAPI.Update)
	http.HandleFunc("GET /area/{id}", areaAPI.Read)

	http.HandleFunc("POST /area/{area_id}/user/{user_id}", areaAPI.AddUser)
	http.HandleFunc("DELETE /area/{area_id}/user/{user_id}", areaAPI.DeleteUser)
	http.HandleFunc("GET /user/{user_id}/list-areas", userAPI.ListAreaByUsers)
	http.HandleFunc("GET /area/{area_id}/list-users", areaAPI.ListUsersByArea)

	http.HandleFunc("POST /user", userAPI.Create)
	http.HandleFunc("PATCH /user/{id}", userAPI.Update)
	http.HandleFunc("GET /user/{id}", userAPI.Read)
	//http.HandleFunc("PATCH /user/area/{id}", userAPI.UpdateUserArea)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}
