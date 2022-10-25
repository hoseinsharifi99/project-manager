package handler

import (
	"amani/auth"
	"amani/db_manager"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type Handler struct {
	dm  *db_manager.DbManager
	ech *echo.Echo
}

func NewHandler(dm *db_manager.DbManager) *Handler {
	h := &Handler{dm: dm, ech: echo.New()}
	h.defineRoutes()
	return h
}

func (h *Handler) defineRoutes() {

	h.ech.Use(auth.JWT())

	auth.AddToWhiteList("/user/login", "POST")
	auth.AddToWhiteList("/user", "POST")

	h.ech.POST("/user", h.SignUp)
	h.ech.POST("/user/login", h.Login)

	h.ech.POST("/prj", h.CreateProject)

	h.ech.POST("/work", h.work)
}

func (h *Handler) Start() {
	h.ech.Logger.Fatal(h.ech.Start(":8080"))
}

func extractID(c echo.Context) uint {
	e := c.Get("user").(*jwt.Token)
	claims := e.Claims.(jwt.MapClaims)
	id := uint(claims["id"].(float64))
	return id
}
