package engine

import (
	"github.com/labstack/echo/v4"
	"reminder/app/api"
)

func NewRoutes() *echo.Echo {
	engine := echo.New()
	// 分组路由
	robotGroup := engine.Group("/robot")
	{
		robotGroup.POST("/add", api.RootApi.CreateRobot)
		robotGroup.POST("/get_all", api.RootApi.GetAllRobots)
		robotGroup.POST("/get_by_page", api.RootApi.GetRobotsByPage)
	}

	userGroup := engine.Group("/user")
	{
		userGroup.POST("/login", api.RootApi.Login)
		userGroup.POST("/add", api.RootApi.CreateUser)
	}
	return engine
}
