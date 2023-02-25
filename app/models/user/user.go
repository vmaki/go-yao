package user

import (
	"go-yao/app/models"
	"go-yao/common/global"
)

// User 用户模型
type User struct {
	models.BaseModel

	Name     string `json:"name,omitempty"`
	Email    string `json:"-"`
	Phone    string `json:"-"`
	Password string `json:"-"`

	models.CommonTimestampsField
}

// Create 创建用户
func (u *User) Create() {
	global.DB.Create(&u)
}
