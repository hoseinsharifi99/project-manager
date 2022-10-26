package handler

import (
	"amani/model"
	"github.com/labstack/echo"
	"net/http"
)

func (h *Handler) CreateProject(c echo.Context) error {
	userID := extractID(c)

	user, err := h.dm.GetUserByID(userID)
	if err != nil {
		echo.NewHTTPError(http.StatusBadRequest, "user doenst exist", err)
	}

	if user.Admin == 1 {
		prj := &model.Project{}
		if err := c.Bind(prj); err != nil {
			echo.NewHTTPError(http.StatusBadRequest, "error binding request", err)
		}

		if err := h.dm.AddProject(prj); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Error adding project to database", err)
		}

		return c.JSON(http.StatusCreated, "project created successfully")
	}

	return c.JSON(http.StatusForbidden, "you are not admin")
}
