package main

import (
	"fmt"
	"projetoinfiel/internal/api"
	"projetoinfiel/internal/database"
	"projetoinfiel/internal/repositories"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	e := echo.New()

	// Adiciona middleware de log e recuperação
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/area", areaAPI.Create)
	e.PATCH("/area/:id", areaAPI.Update)
	e.GET("/area/:id", areaAPI.Read)

	e.POST("/area/:area_id/user/:user_id", areaAPI.AddUser)
	e.DELETE("/area/:area_id/user/:user_id", areaAPI.DeleteUser)
	e.GET("/user/:user_id/list-areas", userAPI.ListAreaByUsers)
	e.GET("/area/:area_id/list-users", areaAPI.ListUsersByArea)

	e.POST("/user", userAPI.Create)
	e.PATCH("/user/:id", userAPI.Update)
	e.GET("/user/:id", userAPI.Read)

	e.Logger.Fatal(e.Start(":8080"))
}
