package handler

import (
	"amani/model"
	"github.com/labstack/echo"
	"net/http"
	"time"
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

func (h *Handler) FetchALlProject(c echo.Context) error {

	projects, err := h.dm.GetProjects()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "error in geting projects from db", err)
	}
	respone := NewProjectsListResponse(projects)
	return c.JSON(http.StatusOK, respone)
}

func NewProjectsListResponse(list []model.Project) *prjListResponse {
	prjs := make([]*prjResponse, 0)
	for i := range list {
		prjs = append(prjs, NewPrjResponse(&list[i]))
	}
	return &prjListResponse{
		Projects: prjs,
		PrjCount: len(prjs),
	}
}

type prjListResponse struct {
	Projects []*prjResponse `json:"projects"`
	PrjCount int            `json:"projects_count"`
}

type prjResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func NewPrjResponse(prj *model.Project) *prjResponse {
	return &prjResponse{
		ID:        int(prj.ID),
		Name:      prj.Name,
		CreatedAt: prj.CreatedAt,
	}
}
