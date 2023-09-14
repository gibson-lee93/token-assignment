package main

import (
	"log"
	"token-assignment/internal/database"
	_pHandler "token-assignment/internal/project/handler"
	_pRepository "token-assignment/internal/project/repository"
	_pService "token-assignment/internal/project/service"

	"github.com/labstack/echo/v4"
)

func main() {
	log.Println("Server Start")

	projectDB := database.GetMySqlDatabase()
	pRepository := _pRepository.NewProjectRepository(projectDB)
	pService := _pService.NewProjectService(pRepository)

	e := echo.New()

	_pHandler.NewProjectHandler(e, pService)

	e.Logger.Fatal(e.Start(":1323"))
}
