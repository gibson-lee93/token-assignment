package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *ProjectHandler) HealthCheck(c echo.Context) error {
	log.Println("Health Check API Called")
	return c.NoContent(http.StatusOK)
}
