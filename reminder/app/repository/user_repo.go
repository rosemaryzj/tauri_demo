package repository

import (
	"go.uber.org/zap"
	"reminder/app/entry"
	"reminder/internal/tools"
)

// CreateUser
// @Description: create user
func (rr *RootRepo) CreateUser(user *entry.User) error {
	user.ID = tools.GetUUID()
	return rr.pool.Create(user).Error
}

// UpdateUser
// @Description: update user info
func (rr *RootRepo) UpdateUser(user entry.User) error {
	return rr.pool.Where("username = ?", user.Username).
		Updates(user).Error
}

// GetUserByName
// @Description: get user info by username
func (rr *RootRepo) GetUserByName(username string) (*entry.User, error) {
	var (
		err  error
		user entry.User
	)
	if err = rr.pool.Where("username = ?", username).
		First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// CheckUser
// @Description: validate user
func (rr *RootRepo) CheckUser(username, password string) (bool, error) {
	var (
		err   error
		count int64
	)
	if err = rr.pool.Model(new(entry.User)).Where("username = ? AND password = ?", username, password).
		Count(&count).Error; err != nil || count == 0 {
		zap.L().Sugar().Errorf("query failed, err: %v", err)
		return false, err
	}
	return true, nil
}

// GetUsersByPageRepo
// @Description get users by pagination
func (rr *RootRepo) GetUsersByPageRepo(page, pageSize int) ([]*entry.User, *int64, error) {
	var (
		err   error
		count int64
		users []*entry.User
	)
	offset := (page - 1) * pageSize
	if err = rr.pool.Model(new(entry.User)).Count(&count).
		Offset(offset).
		Limit(pageSize).
		Find(&users).Error; err != nil {
		return nil, nil, err
	}
	return users, &count, nil
}
