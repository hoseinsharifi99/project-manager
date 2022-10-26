package handler

import (
	"amani/model"
	"github.com/labstack/echo"
	"net/http"
)

func (h *Handler) SignUpAdmin(c echo.Context) error {
	authRequest, err := bindToAuthRequest(c)
	if err != nil {
		return err
	}

	user := &model.User{
		Username: authRequest.Username,
		Password: authRequest.Password,
		Admin:    1,
	}

	user.Password, _ = model.HashPassword(user.Password)
	err = h.dm.AddUser(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "could not add user to database", err)
	}

	return c.JSON(http.StatusCreated, CreateResponseUser(user))
}
