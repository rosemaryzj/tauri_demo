package repository

import (
	"gorm.io/gorm"
	"reminder/app/entry"
	"reminder/thirdparty/lite"
)

// RootRepo
// @Param RobotRepoInterface: database layer interface
// @Param wire: database layer dependencies
type RootRepo struct {
	pool *gorm.DB
	UserRepoInterface
	RobotRepoInterface
}

func NewRootRepo() *RootRepo {
	return &RootRepo{
		pool: lite.NewSqlLite(),
	}
}

// RobotRepoInterface
// @Description interface definition for robot repository api
type RobotRepoInterface interface {
	CreateRobotRepo(robot *entry.Robot) error
	DeleteRobotByNameRepo(robotName string) error
	GetAllRobotsRepo() (robots []*entry.Robot, total *int64, err error)
	UpdateRobotByNameRepo(robotName string, newRobot entry.Robot) error
	GetRobotsByPageRepo(page, pageSize int) ([]*entry.Robot, *int64, error)
}

// UserRepoInterface
// @Description interface definition for user repository api
type UserRepoInterface interface {
	UpdateUser(user entry.User) error
	CreateUserRepo(user *entry.User) error
	GetUserByName(username string) (*entry.User, error)
}
