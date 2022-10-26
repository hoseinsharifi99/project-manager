package handler

import (
	"amani/model"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

type taskReq struct {
	Projectname string `json:"project"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (h *Handler) CreateTask(c echo.Context) error {

	reqTask, err := bindToTaskCreateRequest(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, " invalid prj", err)
	}

	prj, err := h.dm.GetProjectByName(reqTask.Projectname)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "cant find project")
	}

	task := &model.Task{
		ProjectID:   prj.ID,
		Name:        reqTask.Name,
		Duration:    0,
		Description: reqTask.Description,
	}

	fmt.Println("nameeeeeeeeeeee", task.Name)
	fmt.Println("nameeeeeeeeeeee---------", task.ProjectID)

	if err := h.dm.AddTask(task); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Error adding task to database", err)
	}

	return c.JSON(http.StatusCreated, "task created successfully")

}

func bindToTaskCreateRequest(c echo.Context) (*taskReq, error) {
	reqTask := &taskReq{}
	if err := c.Bind(reqTask); err != nil {
		echo.NewHTTPError(http.StatusBadRequest, "error binding request", err)
	}

	return reqTask, nil
}
