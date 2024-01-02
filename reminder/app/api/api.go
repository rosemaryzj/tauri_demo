package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"reminder/app/entry"
	"reminder/app/service"
	"reminder/internal/e"
	"reminder/internal/giveback"
)

var RootApi *RootController

func init() {
	newRootController()
}

// RootController
// @Description service handler
// @Param echo.Context: interface of echo framework
type RootController struct {
	echo.Context
	rootService *service.RootService
}

func newRootController() {
	RootApi = &RootController{
		rootService: service.NewRootService(),
	}
}

// ServeJson
// @Description binding method for self defined controller
// @Param code: http code
// @Param data: json data type
func (rc *RootController) ServeJson(c echo.Context, code int, total *int64, message string, data any) error {
	response := new(giveback.RobotResponse)
	response.Data = data
	response.Total = total
	response.Meta.Code = code
	response.Meta.Message = message
	return c.JSON(http.StatusOK, response)
}

type createRobotParam struct {
	ID           string `json:"id"`
	SN           string `json:"sn"`
	Name         string `json:"name"`
	Username     string `json:"username"`
	InternalIP   string `json:"internal_ip"`
	ExternalIP   string `json:"external_ip"`
	BorrowedUser string `json:"borrowed_user"`
	BorrowedUsed string `json:"borrowed_used"`
	BorrowedDate string `json:"borrowed_date"`
}

// CreateRobot
// @Description create new robot record
func (rc *RootController) CreateRobot(c echo.Context) (err error) {
	param := new(createRobotParam)
	if err = c.Bind(param); err != nil {
		return rc.ServeJson(c, e.JsonBindingErrorCode, nil, e.JsonBindingError, nil)
	}
	if err = rc.rootService.CreateRobotService(param.Username, &entry.Robot{
		SN:           param.SN,
		Name:         param.Name,
		InternalIP:   param.InternalIP,
		ExternalIP:   param.ExternalIP,
		BorrowedUser: param.BorrowedUser,
		BorrowedUsed: param.BorrowedUsed,
		BorrowedDate: param.BorrowedDate,
	}); err != nil {
		return rc.ServeJson(c, e.CreateRobotErrorCode, nil, err.Error(), nil)
	}
	return rc.ServeJson(c, e.SucceedCode, nil, e.SucceedMessage, nil)
}

type PaginationParam struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

// GetUsersByPage
// @Description: get users with pagination
func (rc *RootController) GetUsersByPage(c echo.Context) error {
	var (
		err   error
		total *int64
		users []*entry.User
	)
	param := new(PaginationParam)
	if err = c.Bind(param); err != nil {
		return rc.ServeJson(c, e.JsonBindingErrorCode, nil, e.JsonBindingError, nil)
	}
	if users, total, err = rc.rootService.GetUsersByPageService(param.Page, param.PageSize); err != nil {
		return rc.ServeJson(c, e.GetUsersByPageErrorCode, nil, err.Error(), nil)
	}
	return rc.ServeJson(c, e.SucceedCode, total, e.SucceedMessage, users)
}

// GetRobotsByPage
// @Description: get robot with pagination
func (rc *RootController) GetRobotsByPage(c echo.Context) error {
	var (
		err    error
		total  *int64
		robots []*entry.Robot
	)
	param := new(PaginationParam)
	if err = c.Bind(param); err != nil {
		return rc.ServeJson(c, e.JsonBindingErrorCode, nil, e.JsonBindingError, nil)
	}
	if robots, total, err = rc.rootService.GetRobotsByPageService(param.Page, param.PageSize); err != nil {
		return rc.ServeJson(c, e.GetRobotsByPageErrorCode, nil, err.Error(), nil)
	}
	return rc.ServeJson(c, e.SucceedCode, total, e.SucceedMessage, robots)
}

// GetAllRobots
// @Description: get robot without pagination
func (rc *RootController) GetAllRobots(c echo.Context) (err error) {
	robots, total, err := rc.rootService.GetAllRobotsService()
	if err != nil {
		return rc.ServeJson(c, e.GetAllRobotsErrorCode, nil, err.Error(), nil)
	}
	return rc.ServeJson(c, e.SucceedCode, total, e.SucceedMessage, robots)
}

// CreateUser
// @Description: create user controller
func (rc *RootController) CreateUser(c echo.Context) (err error) {
	var user entry.User
	if err = c.Bind(&user); err != nil {
		return rc.ServeJson(c, e.JsonBindingErrorCode, nil, e.JsonBindingError, nil)
	}

	if err = rc.rootService.CreateUserService(&user); err != nil {
		return rc.ServeJson(c, e.CreateRobotErrorCode, nil, err.Error(), nil)
	}
	return rc.ServeJson(c, e.SucceedCode, nil, e.SucceedMessage, nil)
}

type LoginParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login
// @Description: user login
func (rc *RootController) Login(c echo.Context) error {
	var (
		err   error
		param LoginParam
	)
	if err = c.Bind(&param); err != nil {
		return rc.ServeJson(c, e.JsonBindingErrorCode, nil, e.JsonBindingError, nil)
	}
	tk, err := rc.rootService.LoginService(param.Username, param.Password)
	if err != nil {
		return rc.ServeJson(c, e.GetTokenErrorCode, nil, e.GetTokenError, nil)
	}
	return rc.ServeJson(c, e.SucceedCode, nil, e.SucceedMessage, tk)
}
