package handler

import (
	"encoding/json"
	"my-trophy-prototype-api/domain/model"
	"my-trophy-prototype-api/domain/model/response"
	"my-trophy-prototype-api/interface/database"
	"my-trophy-prototype-api/modules"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) FindTrophies(c echo.Context) error {
	errorResponse := func(code int, title string) error {
		problem := response.Problem{
			Type:  "https://github.com/munierujp/my-trophy-prototype-api/blob/master/handler/trophy.go",
			Title: title,
		}
		c.Response().Header().Set(echo.HeaderContentType, "application/problem+json")
		c.Response().WriteHeader(code)
		return json.NewEncoder(c.Response()).Encode(problem)
	}

	query := &model.Trophy{}
	userID := c.QueryParam("user_id")
	if userID != "" {
		if userID, err := modules.Atouint(userID); err != nil {
			return errorResponse(http.StatusBadRequest, "Invalid User ID")
		} else {
			query.UserID = userID
		}
	}

	trophyRepo := database.NewTrophyRepository(h.DB)
	trophies, err := trophyRepo.Find(query)
	if err != nil {
		return errorResponse(http.StatusBadRequest, "Not found trophy")
	}

	return c.JSON(http.StatusOK, trophies)
}
