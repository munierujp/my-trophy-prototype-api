package handler

import (
	"encoding/json"
	"my-trophy-prototype-api/domain/model/response"
	"my-trophy-prototype-api/interface/database"
	"my-trophy-prototype-api/modules"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) FindUserByID(c echo.Context) error {
	errorResponse := func(code int, title string) error {
		problem := response.Problem{
			Type:  "https://github.com/munierujp/my-trophy-prototype-api/blob/master/handler/user.go",
			Title: title,
		}
		c.Response().Header().Set(echo.HeaderContentType, "application/problem+json")
		c.Response().WriteHeader(code)
		return json.NewEncoder(c.Response()).Encode(problem)
	}

	id, err := modules.Atouint(c.Param("id"))
	if err != nil {
		return errorResponse(http.StatusBadRequest, "Invalid ID")
	}

	userRepo := database.NewUserRepository(h.DB)
	user, err := userRepo.FindByID(id)
	if err != nil {
		return errorResponse(http.StatusBadRequest, "Not found user")
	}

	return c.JSON(http.StatusOK, user)
}