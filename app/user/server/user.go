package server

import (
	"context"
	"fmt"
	"micro_shopping/app/user/dao"
	"micro_shopping/app/user/dao/model"
	"micro_shopping/idl/pb"
	"micro_shopping/pkg/utils"
)

type UserService struct {
}

func (usv *UserService) UserRegister(ctx context.Context, req *pb.UserRequest) (res *pb.UserResponse, err error) {
	res = new(pb.UserResponse)
	daoUser := dao.NewUserDao(ctx)
	//判断用户是否存在
	_, err = daoUser.FindUserByName(req.UserName)
	if err == nil {
		fmt.Println("用户已经存在，请重新输入用户名", err)
		res.Code = 500
		return res, err
	}
	//判断两次输入密码是否相同
	if req.Password != req.PasswordConfirm {
		fmt.Println("两次密码不相同", err)
		res.Code = 500
		return res, err
	}
	//创建用户
	user := model.NewUser(req.UserName, req.Password)
	//加salt和密码加密
	user.CreateSalt()
	user.CreateHashPassword()
	//数据库创建用户
	err = daoUser.CreateUser(user)
	if err != nil {
		fmt.Println("创建用户失败", err)
		res.Code = 500
		return res, err
	}
	res.Code = 200
	res.UserDetail = BuildUserModel(user)
	return res, nil
}

func (usv *UserService) UserLogin(ctx context.Context, req *pb.UserRequest) (res *pb.UserResponse, err error) {
	res = new(pb.UserResponse)
	daoUser := dao.NewUserDao(ctx)
	// 判断用户是否存在
	user, err := daoUser.FindUserByName(req.UserName)
	if err != nil {
		fmt.Println("用户名不存在", err)
		res.Code = 500
		return res, err
	}
	// 验证密码是否正确
	ok := user.CheckPassword(req.Password)
	if !ok {
		fmt.Println("密码错误", err)
		res.Code = 400
		return res, err
	}
	// get claims
	claims := utils.ParseToken(user.Token, utils.MySecret)
	if claims == nil {
		user.Token, _ = utils.CreateToken(req.UserName, user.ID)
		err = daoUser.UpdateUser(user)
		if err != nil {
			fmt.Println("update user token error", err)
			res.Code = 500
			return res, nil
		}
	}
	res.Code = 200
	res.UserDetail = BuildUserModel(user)
	return res, nil
}

func BuildUserModel(user *model.User) *pb.UserModel {
	return &pb.UserModel{
		Id:       uint32(user.ID),
		UserName: user.UserName,
		Token:    user.Token,
	}
}
