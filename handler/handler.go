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
	auth.AddToWhiteList("/admin", "POST")

	//user signup and login
	h.ech.POST("/user", h.SignUp)
	h.ech.POST("/user/login", h.Login)

	//admin signup
	h.ech.POST("/admin", h.SignUpAdmin)

	//create task and project
	h.ech.POST("/prj", h.CreateProject)
	h.ech.POST("/task", h.CreateTask)

	//do task
	h.ech.POST("/work", h.work)

	//get all project
	h.ech.GET("/prj", h.FetchALlProject)

	//get task that user do
	h.ech.GET("/user/task", h.UserTasks)
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
