package database

import (
	"database/sql"
	"go-yao/pkg/global"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var SqlDB *sql.DB

func Connect(dialector gorm.Dialector, log logger.Interface) {
	var err error

	global.DB, err = gorm.Open(dialector, &gorm.Config{Logger: log})
	if err != nil {
		panic("Database connection failure, err: " + err.Error())
	}

	SqlDB, err = global.DB.DB()
	if err != nil {
		panic("SqlDB retrieval failure, err: " + err.Error())
	}
}
