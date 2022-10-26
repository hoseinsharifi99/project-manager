package handler

import (
	"amani/model"
	"github.com/labstack/echo"
	"net/http"
)

type taskReq struct {
	projectname string `json:"project"`
	name        string `json:"name"`
	description string `json:"description"`
}

func (h *Handler) CreateTask(c echo.Context) error {

	reqTask := &taskReq{}
	if err := c.Bind(taskReq{}); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "error binding request", err)
	}

	prj, err := h.dm.GetProjectByName(reqTask.projectname)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "cant find project")
	}
	task := &model.Task{
		ProjectID:   prj.ID,
		Name:        reqTask.name,
		Duration:    0,
		Description: reqTask.description,
	}

	if err := h.dm.AddTask(task); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Error adding task to database", err)
	}

	return c.JSON(http.StatusCreated, "task created successfully")

}
