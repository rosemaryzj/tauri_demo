package service

import "reminder/app/entry"

type RobotServiceInterface interface {
	CreateRobotService(robot *entry.Robot) error
	DeleteRobotByNameService(robotName string) error
	GetRobotsByPageService(page, pageSize int) ([]*entry.Robot, *int64, error)
	UpdateRobotByNameService(robotName string, robot entry.Robot) error
	GetAllRobotsService() (robots []*entry.Robot, total *int64, err error)
}
