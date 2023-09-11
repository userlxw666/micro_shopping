package dao

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"micro_shopping/app/user/dao/model"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(ctx context.Context) *UserDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &UserDao{NewSqlClient(ctx)}
}

func (dao *UserDao) FindUserByName(username string) (*model.User, error) {
	var user *model.User
	err := dao.Where("user_name=?", username).First(&user).Error
	if err != nil {
		fmt.Println("find user error", err)
		return nil, err
	}
	return user, nil
}

func (dao *UserDao) CreateUser(user *model.User) error {
	err := dao.Create(&user).Error
	if err != nil {
		fmt.Println("create user error", err)
		return err
	}
	return nil
}

func (dao *UserDao) UpdateUser(user *model.User) error {
	err := dao.Save(&user).Error
	if err != nil {
		fmt.Println("update user error", err)
		return err
	}
	return nil
}
