package service

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"reminder/app/entry"
	"reminder/app/repository"
	"reminder/internal/tools"
	"reminder/thirdparty/token"
	"time"
)

type RootService struct {
	rootRepo *repository.RootRepo
}

func NewRootService() *RootService {
	return &RootService{
		rootRepo: repository.NewRootRepo(),
	}
}

func (rs *RootService) LoginService(username, password string) (string, error) {
	var (
		err     error
		isAdmin bool
	)
	if username == "admin" {
		isAdmin = true
	}

	exist, err := rs.rootRepo.CheckUser(username, password)
	if err != nil {
		return "", errors.New("sql query error")
	}
	if !exist {
		return "", errors.New("user not found")
	}

	claim := &token.JwtTokenClaim{
		Admin:    isAdmin,
		Username: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	tk := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	t, err := tk.SignedString([]byte(username))
	if err != nil {
		return "", err
	}
	return t, nil
}

func (rs *RootService) CreateUserService(user *entry.User) error {
	if user.Username == "" {
		return errors.New("username is missed, please check")
	}
	if user.Password == "" {
		return errors.New("password is missed, please check")
	}
	return rs.rootRepo.CreateUser(user)
}

// GetUsersByPageService
// @Description: get user service by pagination
func (rs *RootService) GetUsersByPageService(page, pageSize int) ([]*entry.User, *int64, error) {
	if page <= 0 || pageSize >= 100 || pageSize <= 0 {
		return nil, nil, errors.New("100 items per page limited or illegal page and page size,please check")
	}
	return rs.rootRepo.GetUsersByPageRepo(page, pageSize)
}

func (rs *RootService) CreateRobotService(username string, robot *entry.Robot) error {
	var (
		err  error
		user *entry.User
	)
	if robot.Name == "" {
		return errors.New("robot name is missed,please check")
	}
	if robot.SN == "" {
		return errors.New("robot sn is missed, please check")
	}
	if robot.BorrowedUser == "" {
		return errors.New("robot borrow user is missed, please check")
	}
	if robot.ExternalIP == "" {
		return errors.New("robot external ip is missed, please check")
	}
	if robot.InternalIP == "" {
		return errors.New("robot internal ip is missed, please check")
	}
	if robot.BorrowedDate == "" {
		return errors.New("robot borrow date is missed, please check")
	}
	if !tools.IsValidDateTime(robot.BorrowedDate) {
		return errors.New("illegal robot borrow date format, please check")
	}
	// update user related robots count
	if user, err = rs.rootRepo.GetUserByName(username); err != nil {
		return err
	}
	user.RobotCount += 1
	if err = rs.rootRepo.UpdateUser(*user); err != nil {
		return err
	}
	return rs.rootRepo.CreateRobotRepo(robot)
}

func (rs *RootService) UpdateRobotByNameService(robotName string, robot entry.Robot) error {
	if robotName == "" {
		return errors.New("robot name not passwd,please check")
	}
	if robot.BorrowedDate != "" && !tools.IsValidDateTime(robot.BorrowedDate) {
		return errors.New("illegal robot borrow date format, please check")
	}
	return rs.rootRepo.UpdateRobotByNameRepo(robotName, robot)
}

func (rs *RootService) DeleteRobotByNameService(robotName string) error {
	if robotName == "" {
		return errors.New("robot name not passwd,please check")
	}
	return rs.rootRepo.DeleteRobotByNameRepo(robotName)
}

func (rs *RootService) GetRobotsByPageService(page, pageSize int) ([]*entry.Robot, *int64, error) {
	if page <= 0 || pageSize >= 100 || pageSize <= 0 {
		return nil, nil, errors.New("100 items per page limited or illegal page and page size,please check")
	}
	return rs.rootRepo.GetRobotsByPageRepo(page, pageSize)
}

func (rs *RootService) GetAllRobotsService() ([]*entry.Robot, *int64, error) {
	return rs.rootRepo.GetAllRobotsRepo()
}
