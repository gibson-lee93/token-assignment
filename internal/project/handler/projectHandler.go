package handler

import (
	"token-assignment/internal/project/service"

	"github.com/labstack/echo/v4"
)

type ProjectHandler struct {
	svc service.ProjectService
}

func NewProjectHandler(e *echo.Echo, svc service.ProjectService) {
	h := &ProjectHandler{svc: svc}
	{
		e.GET("/healthcheck", h.HealthCheck)
	}
}
