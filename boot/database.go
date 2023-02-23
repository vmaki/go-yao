package boot

import (
	"errors"
	"fmt"
	"go-yao/common/global"
	"go-yao/pkg/database"
	"go-yao/pkg/logger"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// SetupDB 初始化数据库和 ORM
func SetupDB() {
	config := global.Conf.DataBase

	var dialector gorm.Dialector

	// 根据情况可以加载对应驱动
	switch config.Driver {
	case "mysql":
		dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&multiStatements=true&loc=Local",
			config.Username,
			config.Password,
			config.Host,
			config.Port,
			config.Database,
			config.Charset,
		)
		dialector = mysql.New(mysql.Config{
			DSN: dsn,
		})
	default:
		panic(errors.New("database connection not supported"))
	}

	// 连接数据库，并设置 GORM 的日志模式
	database.Connect(dialector, logger.NewGormLogger())

	database.SqlDB.SetMaxOpenConns(config.MaxOpenConnections)
	database.SqlDB.SetMaxIdleConns(config.MaxIdleConnections)
	database.SqlDB.SetConnMaxLifetime(time.Duration(config.MaxLifeSeconds) * time.Second)

	// global.DB.AutoMigrate(&model.User{}) // 自动迁移
}
