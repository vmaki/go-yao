package user

import (
	"github.com/gin-gonic/gin"
	"go-yao/common/global"
	"go-yao/pkg/paginator"
)

// IsPhoneExist 判断手机号已被注册
func IsPhoneExist(phone string) bool {
	var count int64
	global.DB.Model(User{}).Where("phone = ?", phone).Count(&count)
	return count > 0
}

// GetByPhone 通过手机号来获取用户
func GetByPhone(phone string) (userModel User) {
	global.DB.Where("phone = ?", phone).First(&userModel)
	return
}

func Paginate(ctx *gin.Context, perPage int) (users []User, paging paginator.Paging) {
	paging = paginator.Paginate(
		ctx,
		global.DB.Model(User{}),
		&users,
		perPage,
	)

	return
}
