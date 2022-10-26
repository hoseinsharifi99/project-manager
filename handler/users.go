package handler

import (
	"amani/auth"
	"amani/model"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

type userAuthRequest struct {
	Username string `json:"user_name"`
	Password string `json:"password"`
}

func bindToAuthRequest(c echo.Context) (*userAuthRequest, error) {
	var userAuth = &userAuthRequest{}
	if err := c.Bind(userAuth); err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "error binding user", err)
	}
	return userAuth, nil
}

func (h *Handler) SignUp(c echo.Context) error {
	authRequest, err := bindToAuthRequest(c)
	if err != nil {
		return err
	}

	user := &model.User{
		Username: authRequest.Username,
		Password: authRequest.Password,
	}

	user.Password, _ = model.HashPassword(user.Password)
	err = h.dm.AddUser(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "could not add user to database", err)
	}

	return c.JSON(http.StatusCreated, CreateResponseUser(user))
}

type UserResponse struct {
	UserName string `json:"user_name"`
	Token    string `json:"token"`
}

func CreateResponseUser(user *model.User) *UserResponse {
	token, _ := auth.GenerateJWT(user.ID)
	resUser := &UserResponse{
		UserName: user.Username,
		Token:    token,
	}
	return resUser
}

func (h *Handler) Login(c echo.Context) error {
	authRequest, err := bindToAuthRequest(c)
	if err != nil {
		return err
	}

	user := &model.User{
		Username: authRequest.Username,
		Password: authRequest.Password,
	}
	u, err := h.dm.GetUserByUsername(user.Username)
	if err != nil || !u.ValidatePassword(user.Password) {
		return echo.NewHTTPError(http.StatusBadRequest, "invalide username or password", err)
	}
	return c.JSON(http.StatusOK, CreateResponseUser(u))
}

func (h *Handler) UserTasks(c echo.Context) error {
	userID := extractID(c)
	userProject, err := h.dm.GetTaskByUserID(userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Error retrieving usertasks from database", err)
	}
	response := h.NewUserTasksListResponse(userProject)
	return c.JSON(http.StatusOK, response)
}

type userTaskResponse struct {
	ID          int       `json:"id"`
	UserName    string    `json:"username"`
	ProjectName string    `json:"project"`
	TaskName    string    `json:"task"`
	CreatedAt   time.Time `json:"created_at"`
	LastUpdate  time.Time `json:"last_update"`
	Duration    int       `json:"duration"`
}

func NewtaskResponse(prj *model.Project, task *model.Task, up *model.UserProject, user *model.User) *userTaskResponse {
	return &userTaskResponse{
		ID:          int(up.ID),
		UserName:    user.Username,
		ProjectName: prj.Name,
		TaskName:    task.Name,
		CreatedAt:   up.CreatedAt,
		LastUpdate:  up.UpdatedAt,
		Duration:    int(up.Duration),
	}
}

func (h *Handler) NewUserTasksListResponse(list []model.UserProject) *userTaskListResponse {
	prjs := make([]*userTaskResponse, 0)
	for i, uprj := range list {
		user, err := h.dm.GetUserByID(uprj.UserID)
		if err != nil {
			return nil
		}
		prj, err := h.dm.GetProjectById(uprj.ProjectID)
		if err != nil {
			return nil
		}
		tsk, err := h.dm.GetTaskById(uprj.TaskID)
		if err != nil {
			return nil
		}
		prjs = append(prjs, NewtaskResponse(prj, tsk, &list[i], user))
	}
	return &userTaskListResponse{
		Projects: prjs,
		PrjCount: len(prjs),
	}
}

type userTaskListResponse struct {
	Projects []*userTaskResponse `json:"user-tasks"`
	PrjCount int                 `json:"tasks_count"`
}
