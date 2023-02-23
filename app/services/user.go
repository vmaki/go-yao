package services

import (
	"fmt"
	"go-yao/app/model/user"
	"go-yao/common/response"
)

type UserService struct {
}

// Register 注册
func (s *UserService) Register(phone string) (*user.User, error) {
	if isExist := user.IsPhoneExist(phone); isExist {
		return nil, response.New(response.CodeUserExist)
	}

	data := &user.User{
		Name:  s.maskPhone(phone),
		Phone: phone,
	}
	data.Create()

	if data.ID > 0 {
		return data, nil
	}

	return nil, response.New(response.CodeSysError)
}

// maskPhone 隐藏用户手机号码
func (s *UserService) maskPhone(phone string) string {
	if len(phone) < 10 {
		return phone
	}

	return fmt.Sprintf("%s****%s", phone[:3], phone[len(phone)-4:])
}

// LoginByPhone 使用手机号码登录
func (s *UserService) LoginByPhone(phone string) (user.User, error) {
	userModel := user.GetByPhone(phone)
	if userModel.ID == 0 {
		return user.User{}, response.New(response.CodeUserNotExist)
	}

	return userModel, nil
}
