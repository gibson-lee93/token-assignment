package main

import (
	"log"
	"net/http"
	"token-assignment/internal/database"

	"github.com/labstack/echo/v4"
)

func main() {
	log.Println("Server Start")

	database.GetMySqlDatabase()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
