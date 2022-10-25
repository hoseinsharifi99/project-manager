package handler

import (
	"amani/model"
	"github.com/labstack/echo"
	"log"
	"net/http"
)

func (h *Handler) work(c echo.Context) error {
	userID := extractID(c)

	req, err := bindToUrlCreateRequest(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, " invalid prj", err)
	}

	var prj *model.Project
	prj, err = h.dm.GetProjectByName(req.Name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "prj doesnt exit", err)
	}
	var userprj *model.UserProject
	userprj, err = h.dm.GetUserProjects(userID, prj.ID)
	log.Println("in yarooooooooooooooooooooooooooooooooo")
	log.Println(err)
	if err != nil {
		userProject := &model.UserProject{
			UserID:    userID,
			ProjectID: prj.ID,
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
	Name     string `json:"name"`
	Duration uint   `json:"duration"`
}

func bindToUrlCreateRequest(c echo.Context) (*userProjectRequest, error) {
	request := &userProjectRequest{}
	if err := c.Bind(request); err != nil {
		echo.NewHTTPError(http.StatusBadRequest, "error binding request", err)
	}

	return request, nil
}
