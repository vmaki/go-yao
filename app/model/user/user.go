package user

import (
	"go-yao/app/model"
	"go-yao/common/global"
)

// User 用户模型
type User struct {
	model.BaseModel

	Name     string `json:"name,omitempty"`
	Email    string `json:"-"`
	Phone    string `json:"-"`
	Password string `json:"-"`

	model.CommonTimestampsField
}

// Create 创建用户
func (u *User) Create() {
	global.DB.Create(&u)
}
