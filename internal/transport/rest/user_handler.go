package rest

import (
	"net/http"

	"github.com/banggibima/go-echo-restful-api/internal/entities"
	"github.com/banggibima/go-echo-restful-api/internal/usecases"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UserUseCase *usecases.UserUseCase
}

func NewUserHandler(userUseCase *usecases.UserUseCase) *UserHandler {
	return &UserHandler{UserUseCase: userUseCase}
}

func (h *UserHandler) GetUsersHandler(c echo.Context) error {
	users, err := h.UserUseCase.GetUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, users)
}

func (h *UserHandler) GetUserByIDHandler(c echo.Context) error {
	id := c.Param("id")

	user, err := h.UserUseCase.GetUserByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) CreateUserHandler(c echo.Context) error {
	var input entities.User

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	user, err := h.UserUseCase.CreateUser(input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) UpdateUserHandler(c echo.Context) error {
	id := c.Param("id")

	var input entities.User

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	user, err := h.UserUseCase.UpdateUser(id, input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) DeleteUserHandler(c echo.Context) error {
	id := c.Param("id")

	if err := h.UserUseCase.DeleteUser(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}
