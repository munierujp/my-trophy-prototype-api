package handler

import (
	"encoding/json"
	"my-trophy-prototype-api/domain/model"
	"my-trophy-prototype-api/domain/model/response"
	"my-trophy-prototype-api/interface/database"
	"my-trophy-prototype-api/modules"
	"net/http"
	"unicode/utf8"

	"firebase.google.com/go/auth"
	"github.com/labstack/echo/v4"
)

func (h *Handler) CreateUser(c echo.Context) error {
	errorResponse := func(code int, title string) error {
		problem := response.Problem{
			Type:  "https://github.com/munierujp/my-trophy-prototype-api/blob/master/handler/user.go",
			Title: title,
		}
		c.Response().Header().Set(echo.HeaderContentType, "application/problem+json")
		c.Response().WriteHeader(code)
		return json.NewEncoder(c.Response()).Encode(problem)
	}

	jwt := c.Get("jwt").(*auth.Token)
	claims := jwt.Claims
	name, ok := claims["name"].(string)
	if !ok {
		return errorResponse(http.StatusBadRequest, "Invalid token")
	}
	if name == "" || utf8.RuneCountInString(name) > 50 {
		return errorResponse(http.StatusBadRequest, "Invalid request")
	}
	email, ok := claims["email"].(string)
	if !ok {
		return errorResponse(http.StatusBadRequest, "Invalid token")
	}
	if email == "" || utf8.RuneCountInString(email) > 256 {
		return errorResponse(http.StatusBadRequest, "Invalid request")
	}

	userRepo := database.NewUserRepository(h.DB)
	query := &model.User{Email: email}
	users, _ := userRepo.Find(query)
	if len(users) > 0 {
		return errorResponse(http.StatusBadRequest, "Already exists")
	}

	user := &model.User{
		Name:  name,
		Email: email,
	}

	if err := userRepo.Create(user); err != nil {
		return errorResponse(http.StatusBadRequest, "Failed to create")
	}

	return c.NoContent(http.StatusCreated)
}

func (h *Handler) FindUsers(c echo.Context) error {
	errorResponse := func(code int, title string) error {
		problem := response.Problem{
			Type:  "https://github.com/munierujp/my-trophy-prototype-api/blob/master/handler/user.go",
			Title: title,
		}
		c.Response().Header().Set(echo.HeaderContentType, "application/problem+json")
		c.Response().WriteHeader(code)
		return json.NewEncoder(c.Response()).Encode(problem)
	}

	query := &model.User{}
	email := c.QueryParam("email")
	if email != "" {
		query.Email = email
	}

	userRepo := database.NewUserRepository(h.DB)
	users, err := userRepo.Find(query)
	if err != nil {
		return errorResponse(http.StatusBadRequest, "Not found user")
	}

	return c.JSON(http.StatusOK, users)
}

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
