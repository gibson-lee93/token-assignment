package handler

import (
	"token-assignment/internal/project/service"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func echoProvision(e *echo.Echo) {
	e.Use(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Validator = &CustomValidator{validator: validator.New()}
}

type ProjectHandler struct {
	svc service.ProjectService
}

func NewProjectHandler(e *echo.Echo, svc service.ProjectService) {
	echoProvision(e)
	h := &ProjectHandler{svc: svc}
	{
		e.GET("/healthcheck", h.HealthCheck)
		e.GET("/tokeninfo", h.GetTokenInfo)
	}
}
