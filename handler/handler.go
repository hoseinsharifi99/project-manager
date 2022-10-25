package handler

import (
	"amani/db_manager"
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

	h.ech.POST("/user", h.SignUp)
}

func (h *Handler) Start() {
	h.ech.Logger.Fatal(h.ech.Start(":8080"))
}
