package repository

import (
	"reminder/app/entry"
	"reminder/internal/tools"
)

func (rr *RootRepo) CreateRobotRepo(robot *entry.Robot) error {
	robot.ID = tools.GetUUID()
	return rr.pool.Create(robot).Error
}

// DeleteRobotByNameRepo
// @Param robotName: should be unique in database
func (rr *RootRepo) DeleteRobotByNameRepo(robotName string) (err error) {
	return rr.pool.Where("name = ?", robotName).
		Delete(new(entry.Robot)).Error
}

// UpdateRobotByNameRepo
// @Param robotName: should be unique in database
func (rr *RootRepo) UpdateRobotByNameRepo(robotName string, newRobot entry.Robot) error {
	return rr.pool.Where("name = ?", robotName).
		Updates(newRobot).Error
}

// GetRobotsByPageRepo
// @Description: get robot by page
func (rr *RootRepo) GetRobotsByPageRepo(page, pageSize int) ([]*entry.Robot, *int64, error) {
	var (
		err    error
		total  int64
		robots []*entry.Robot
	)
	offset := (page - 1) * pageSize
	if err = rr.pool.
		Model(new(entry.Robot)).
		Count(&total).
		Offset(offset).
		Limit(pageSize).
		Find(&robots).Error; err != nil {
		return nil, nil, err
	}
	return robots, &total, nil
}

// GetAllRobotsRepo
// @Description: not suggested(only used for small amount of data)
func (rr *RootRepo) GetAllRobotsRepo() ([]*entry.Robot, *int64, error) {
	var (
		err    error
		total  int64
		robots []*entry.Robot
	)
	if err = rr.pool.Find(&robots).Count(&total).Error; err != nil {
		return nil, nil, err
	}
	return robots, &total, nil
}
