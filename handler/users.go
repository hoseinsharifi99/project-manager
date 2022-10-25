package handler

import (
	"amani/auth"
	"amani/model"
	"github.com/labstack/echo"
	"net/http"
)

type userAuthRequest struct {
	Username string `json:"user_name"`
	Password string `json:"password"`
}

func bindToAuthRequest(c echo.Context) (*userAuthRequest, error) {
	var userAuth = &userAuthRequest{}
	if err := c.Bind(userAuth); err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "error binding user", err)
	}
	return userAuth, nil
}

func (h *Handler) SignUp(c echo.Context) error {
	authRequest, err := bindToAuthRequest(c)
	if err != nil {
		return err
	}

	user := &model.User{
		Username: authRequest.Username,
		Password: authRequest.Password,
	}

	user.Password, _ = model.HashPassword(user.Password)
	err = h.dm.AddUser(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "could not add user to database", err)
	}

	return c.JSON(http.StatusCreated, CreateResponseUser(user))
}

type UserResponse struct {
	UserName string `json:"user_name"`
	Token    string `json:"token"`
}

func CreateResponseUser(user *model.User) *UserResponse {
	token, _ := auth.GenerateJWT(user.ID)
	resUser := &UserResponse{
		UserName: user.Username,
		Token:    token,
	}
	return resUser
}

func (h *Handler) Login(c echo.Context) error {
	authRequest, err := bindToAuthRequest(c)
	if err != nil {
		return err
	}

	user := &model.User{
		Username: authRequest.Username,
		Password: authRequest.Password,
	}
	u, err := h.dm.GetUserByUsername(user.Username)
	if err != nil || !u.ValidatePassword(user.Password) {
		return echo.NewHTTPError(http.StatusBadRequest, "invalide username or password", err)
	}
	return c.JSON(http.StatusOK, CreateResponseUser(u))
}
