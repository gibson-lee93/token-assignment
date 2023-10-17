package handler

import (
	"log"
	"net/http"
	"time"
	"token-assignment/internal/project/dto"

	"github.com/labstack/echo/v4"
)

func (h *ProjectHandler) GetTokenInfo(c echo.Context) error {
	var err error

	symbolStr := c.QueryParam("tokenSymbol")
	sourceStr := c.QueryParam("source")
	startTimeStr := c.QueryParam("startTime")
	endTimeStr := c.QueryParam("endTime")

	dtoReq := dto.GetTokenInfoReq{
		Symbol: symbolStr,
		Source: sourceStr,
	}

	if startTimeStr != "" {
		startTime, err := time.Parse(time.RFC3339, startTimeStr)
		if err != nil {
			return c.JSON(http.StatusOK, dto.Response{ErrorMessage: "Invalid time format"})
		}
		dtoReq.StartTime = startTime

		endTime, err := time.Parse(time.RFC3339, endTimeStr)
		if err != nil {
			return c.JSON(http.StatusOK, dto.Response{ErrorMessage: "Invalid time format"})
		}
		dtoReq.EndTime = endTime
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
