package database

import (
	"database/sql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var SqlDB *sql.DB

func Connect(dialector gorm.Dialector, log logger.Interface) {
	DB, err := gorm.Open(dialector, &gorm.Config{Logger: log})
	if err != nil {
		panic("Database connection failure, err: " + err.Error())
	}

	SqlDB, err = DB.DB()
	if err != nil {
		panic("SqlDB retrieval failure, err: " + err.Error())
	}
}
