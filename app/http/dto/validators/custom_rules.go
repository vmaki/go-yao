package validators

import (
	"errors"
	"fmt"
	"go-yao/pkg/global"
	"strings"

	"github.com/thedevsaddam/govalidator"
)

// 此方法会在初始化时执行，注册自定义表单验证规则
func init() {
	// 自定义规则 not_exists，验证请求数据必须不存在于数据库中。
	// 常用于保证数据库某个字段的值唯一，如用户名、邮箱、手机号、或者分类的名称。
	// not_exists 参数可以有两种，一种是 2 个参数，一种是 3 个参数：
	// not_exists:users,phone 检查数据库表里是否存在同一条信息
	// not_exists:users,phone,32 排除用户掉 id 为 32 的用户
	govalidator.AddCustomRule("not_exists", func(field string, rule string, message string, value interface{}) error {
		rng := strings.Split(strings.TrimPrefix(rule, "not_exists:"), ",")

		tableName := rng[0] // 第一个参数，表名称，如 users
		dbFiled := rng[1]   // 第二个参数，字段名称，如 phone

		// 第三个参数，排除 ID
		var exceptID string
		if len(rng) > 2 {
			exceptID = rng[2]
		}

		requestValue := value.(string)                                          // 用户请求过来的数据
		query := global.DB.Table(tableName).Where(dbFiled+" = ?", requestValue) // 拼接 SQL

		// 如果传参第三个参数，加上 SQL Where 过滤
		if len(exceptID) > 0 {
			query.Where("id != ?", exceptID)
		}

		// 查询数据库
		var count int64
		query.Count(&count)
		if count != 0 {
			// 如果有自定义错误消息的话
			if message != "" {
				return errors.New(message)
			}

			return fmt.Errorf("%v 已存在", requestValue)
		}

		return nil
	})
}
