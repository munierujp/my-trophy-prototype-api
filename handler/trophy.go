package handler

import (
	"encoding/json"
	"my-trophy-prototype-api/domain/model"
	"my-trophy-prototype-api/domain/model/response"
	"my-trophy-prototype-api/interface/database"
	"my-trophy-prototype-api/modules"
	"my-trophy-prototype-api/types/date"
	"net/http"
	"unicode/utf8"

	"firebase.google.com/go/auth"
	"github.com/labstack/echo/v4"
)

type TrophyRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	AchievedOn  string `json:"achieved_on"`
}

func (h *Handler) CreateTrophy(c echo.Context) error {
	errorResponse := func(code int, title string) error {
		problem := response.Problem{
			Type:  "https://github.com/munierujp/my-trophy-prototype-api/blob/master/handler/trophy.go",
			Title: title,
		}
		c.Response().Header().Set(echo.HeaderContentType, "application/problem+json")
		c.Response().WriteHeader(code)
		return json.NewEncoder(c.Response()).Encode(problem)
	}

	jwt := c.Get("jwt").(*auth.Token)
	claims := jwt.Claims
	email, ok := claims["email"].(string)
	if !ok {
		return errorResponse(http.StatusBadRequest, "Invalid token")
	}

	userRepo := database.NewUserRepository(h.DB)
	query := &model.User{Email: email}
	users, err := userRepo.Find(query)
	if err != nil || len(users) == 0 {
		return errorResponse(http.StatusBadRequest, "Not found user")
	}
	user := users[0]

	req := new(TrophyRequest)
	if err := c.Bind(req); err != nil {
		return errorResponse(http.StatusBadRequest, "Invalid request")
	}
	title := req.Title
	if title == "" || utf8.RuneCountInString(title) > 13 {
		return errorResponse(http.StatusBadRequest, "Invalid request")
	}
	description := req.Description
	if utf8.RuneCountInString(description) > 280 {
		return errorResponse(http.StatusBadRequest, "Invalid request")
	}
	if req.AchievedOn == "" {
		return errorResponse(http.StatusBadRequest, "Invalid request")
	}
	achievedOn, err := date.Parse(date.RFC3339, req.AchievedOn)
	if err != nil {
		return errorResponse(http.StatusBadRequest, "Invalid request")
	}

	trophy := &model.Trophy{
		Title:       title,
		Description: description,
		AchievedOn:  achievedOn,
		UserID:      user.ID,
	}

	trophyRepo := database.NewTrophyRepository(h.DB)
	if err := trophyRepo.Create(trophy); err != nil {
		return errorResponse(http.StatusBadRequest, "Failed to create")
	}

	return c.NoContent(http.StatusCreated)
}

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

func (h *Handler) FindTrophyByID(c echo.Context) error {
	errorResponse := func(code int, title string) error {
		problem := response.Problem{
			Type:  "https://github.com/munierujp/my-trophy-prototype-api/blob/master/handler/trophy.go",
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

	trophyRepo := database.NewTrophyRepository(h.DB)
	trophy, err := trophyRepo.FindByID(id)
	if err != nil {
		return errorResponse(http.StatusBadRequest, "Not found trophy")
	}

	return c.JSON(http.StatusOK, trophy)
}

func (h *Handler) UpdateTrophy(c echo.Context) error {
	errorResponse := func(code int, title string) error {
		problem := response.Problem{
			Type:  "https://github.com/munierujp/my-trophy-prototype-api/blob/master/handler/trophy.go",
			Title: title,
		}
		c.Response().Header().Set(echo.HeaderContentType, "application/problem+json")
		c.Response().WriteHeader(code)
		return json.NewEncoder(c.Response()).Encode(problem)
	}

	jwt := c.Get("jwt").(*auth.Token)
	claims := jwt.Claims
	email, ok := claims["email"].(string)
	if !ok {
		return errorResponse(http.StatusBadRequest, "Invalid token")
	}

	id, err := modules.Atouint(c.Param("id"))
	if err != nil {
		return errorResponse(http.StatusBadRequest, "Invalid ID")
	}

	trophyRepo := database.NewTrophyRepository(h.DB)
	trophy, err := trophyRepo.FindByID(id)
	if err != nil {
		return errorResponse(http.StatusBadRequest, "Not found trophy")
	}

	userRepo := database.NewUserRepository(h.DB)
	query := &model.User{Email: email}
	users, err := userRepo.Find(query)
	if err != nil || len(users) == 0 {
		return errorResponse(http.StatusBadRequest, "Not found user")
	}
	user := users[0]

	if trophy.UserID != user.ID {
		return errorResponse(http.StatusBadRequest, "You do'nt have permission")
	}

	req := new(TrophyRequest)
	if err := c.Bind(req); err != nil {
		return errorResponse(http.StatusBadRequest, "Invalid request")
	}
	title := req.Title
	if title == "" {
		return errorResponse(http.StatusBadRequest, "Invalid request")
	}
	description := req.Description
	if req.AchievedOn == "" {
		return errorResponse(http.StatusBadRequest, "Invalid request")
	}
	achievedOn, err := date.Parse(date.RFC3339, req.AchievedOn)
	if err != nil {
		return errorResponse(http.StatusBadRequest, "Invalid request")
	}

	trophy.Title = title
	trophy.Description = description
	trophy.AchievedOn = achievedOn

	if err := trophyRepo.Update(trophy); err != nil {
		return errorResponse(http.StatusBadRequest, "Failed to update")
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *Handler) DeleteTrophy(c echo.Context) error {
	errorResponse := func(code int, title string) error {
		problem := response.Problem{
			Type:  "https://github.com/munierujp/my-trophy-prototype-api/blob/master/handler/trophy.go",
			Title: title,
		}
		c.Response().Header().Set(echo.HeaderContentType, "application/problem+json")
		c.Response().WriteHeader(code)
		return json.NewEncoder(c.Response()).Encode(problem)
	}

	jwt := c.Get("jwt").(*auth.Token)
	claims := jwt.Claims
	email, ok := claims["email"].(string)
	if !ok {
		return errorResponse(http.StatusBadRequest, "Invalid token")
	}

	id, err := modules.Atouint(c.Param("id"))
	if err != nil {
		return errorResponse(http.StatusBadRequest, "Invalid ID")
	}

	trophyRepo := database.NewTrophyRepository(h.DB)
	trophy, err := trophyRepo.FindByID(id)
	if err != nil {
		return errorResponse(http.StatusBadRequest, "Not found trophy")
	}

	userRepo := database.NewUserRepository(h.DB)
	query := &model.User{Email: email}
	users, err := userRepo.Find(query)
	if err != nil || len(users) == 0 {
		return errorResponse(http.StatusBadRequest, "Not found user")
	}
	user := users[0]

	if trophy.UserID != user.ID {
		return errorResponse(http.StatusBadRequest, "You do'nt have permission")
	}

	if err := trophyRepo.Delete(id); err != nil {
		return errorResponse(http.StatusBadRequest, "Failed to delete")
	}

	return c.NoContent(http.StatusNoContent)
}
