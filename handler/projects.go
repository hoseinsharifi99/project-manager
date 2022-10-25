package handler

import (
	"amani/model"
	"github.com/labstack/echo"
	"net/http"
)

func (h *Handler) CreateProject(c echo.Context) error {

	prj := &model.Project{}
	if err := c.Bind(prj); err != nil {
		echo.NewHTTPError(http.StatusBadRequest, "error binding request", err)
	}

	if err := h.dm.AddProject(prj); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Error adding url to database", err)
	}

	return c.JSON(http.StatusCreated, "URL created successfully")
}
