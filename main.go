package main

import (
	"log"
	"time"
	"token-assignment/internal/database"
	_pHandler "token-assignment/internal/project/handler"
	_pRepository "token-assignment/internal/project/repository"
	_pService "token-assignment/internal/project/service"

	"github.com/labstack/echo/v4"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Server Start")

	projectDB := database.GetMySqlDatabase()
	pRepository := _pRepository.NewProjectRepository(projectDB)
	pService := _pService.NewProjectService(pRepository)

	e := echo.New()

	_pHandler.NewProjectHandler(e, pService)

	// scheduler for fetching tiken infos every 30 seconds
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	tokenSymbol := "USTUSD"
	go func() {
		log.Println("Running Scheduler")
		pService.GetBitfinexTokenInfo(tokenSymbol)

		// Loop to handle ticks
		for {
			select {
			case <-ticker.C:
				log.Println("Fetching Token Information")
				pService.GetBitfinexTokenInfo(tokenSymbol)
			}
		}
	}()

	e.Logger.Fatal(e.Start(":1323"))
}
