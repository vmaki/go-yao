package services

import (
	"errors"
	"fmt"
	"go-yao/app/model/user"
)

type UserService struct {
}

// Register 注册
func (s *UserService) Register(phone string) (*user.User, error) {
	if isExist := user.IsPhoneExist(phone); isExist {
		return nil, errors.New("用户已注册")
	}

	data := &user.User{
		Name:  s.maskPhone(phone),
		Phone: phone,
	}
	data.Create()

	if data.ID > 0 {
		return data, nil
	}

	return nil, errors.New("创建用户失败")
}

// maskPhone 隐藏用户手机号码
func (s *UserService) maskPhone(phone string) string {
	if len(phone) < 10 {
		return phone
	}

	return fmt.Sprintf("%s****%s", phone[:3], phone[len(phone)-4:])
}

// LoginByPhone 登录指定用户
func (s *UserService) LoginByPhone(phone string) (user.User, error) {
	userModel := user.GetByPhone(phone)
	if userModel.ID == 0 {
		return user.User{}, errors.New("手机号未注册")
	}

	return userModel, nil
}
