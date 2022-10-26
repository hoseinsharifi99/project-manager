package handler

import (
	"amani/model"
	"github.com/labstack/echo"
	"net/http"
)

type taskReq struct {
	Projectname string `json:"project"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type TasksOfProject struct {
	Name string `json:"name"`
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

	if err := h.dm.AddTask(task); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Error adding task to database", err)
	}

	return c.JSON(http.StatusCreated, "task created successfully")

}

func bindToTaskCreateRequest(c echo.Context) (*taskReq, error) {
	reqTask := &taskReq{}
	if err := c.Bind(reqTask); err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "error binding request", err)
	}

	return reqTask, nil
}

func bindTaskOfProject(c echo.Context) (*TasksOfProject, error) {
	req := &TasksOfProject{}
	if err := c.Bind(req); err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "error binding", err)
	}
	return req, nil
}

func (h *Handler) TaskOfProject(c echo.Context) error {
	prjReq, err := bindTaskOfProject(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, " invalid prj", err)
	}
	tasks, err := h.dm.GetTasksByProjectName(prjReq.Name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, " cant get tasks", err)
	}
	response := NewTaskListResponse(tasks)
	return c.JSON(http.StatusOK, response)
}

type taskResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type taskList struct {
	Tasks     []*taskResponse `json:"tasks"`
	TaskCount int             `json:"taskCount"`
}

func NewTaskResponse(tsk *model.Task) *taskResponse {
	return &taskResponse{
		Name:        tsk.Name,
		Description: tsk.Description,
	}
}

func NewTaskListResponse(list []model.Task) *taskList {
	tasks := make([]*taskResponse, 0)
	for i := range list {
		tasks = append(tasks, NewTaskResponse(&list[i]))
	}
	return &taskList{
		Tasks:     tasks,
		TaskCount: len(tasks),
	}
}
