package handler

import (
	"amani/model"
	"github.com/labstack/echo"
	"log"
	"net/http"
)

func (h *Handler) work(c echo.Context) error {
	userID := extractID(c)

	req, err := bindToUserProjectCreateRequest(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, " invalid prj", err)
	}

	var prj *model.Project
	prj, err = h.dm.GetProjectByName(req.PrjName)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "prj doesnt exit", err)
	}

	var tsk *model.Task
	tsk, err = h.dm.GetTaskByName(req.TaskName)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "task doesnt exit", err)
	}

	var userprj *model.UserProject
	userprj, err = h.dm.GetUserProjects(userID, prj.ID, tsk.ID)

	log.Println(err)
	if err != nil {
		userProject := &model.UserProject{
			UserID:    userID,
			ProjectID: prj.ID,
			TaskID:    tsk.ID,
			Duration:  req.Duration,
		}
		if err := h.dm.AddUserProjec(userProject); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Error adding userprj to database", err)
		}
		return c.JSON(http.StatusCreated, "ok add")
	}

	userprj.Duration += req.Duration
	if err := h.dm.UpdateUserProject(userprj); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Error updating userprj to database", err)
	}

	return c.JSON(http.StatusCreated, "ok update")
}

type userProjectRequest struct {
	PrjName  string `json:"name"`
	TaskName string `json:"task"`
	Duration uint   `json:"duration"`
}

func bindToUserProjectCreateRequest(c echo.Context) (*userProjectRequest, error) {
	request := &userProjectRequest{}
	if err := c.Bind(request); err != nil {
		echo.NewHTTPError(http.StatusBadRequest, "error binding request", err)
	}

	return request, nil
}
