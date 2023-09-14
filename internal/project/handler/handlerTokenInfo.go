package handler

import (
	"log"
	"net/http"
	"token-assignment/internal/project/dto"

	"github.com/labstack/echo/v4"
)

func (h *ProjectHandler) GetTokenInfo(c echo.Context) error {
	var err error
	var dtoReq dto.GetTokenInfoReq

	if err = c.Bind(&dtoReq); err != nil {
		log.Println(err)
		return c.JSON(http.StatusOK, dto.Response{ErrorMessage: "Invalid Body"})
	}

	if err = c.Validate(dtoReq); err != nil {
		log.Println(err)
		return c.JSON(http.StatusOK, dto.Response{ErrorMessage: "Invalid Body"})
	}

	dtoResp, err := h.svc.GetTokenInfo(dtoReq)
	if err != nil {
		return c.JSON(http.StatusOK, dto.Response{ErrorMessage: "Error fetching token info"})
	}

	return c.JSON(http.StatusOK, dtoResp)
}
