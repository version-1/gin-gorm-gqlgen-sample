package pkg

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func Connect(env string) *gorm.DB {
	dbConf := GetConfig(env).Database
	mysqlConf := mysql.Config{
		User:                 dbConf.User,
		Passwd:               dbConf.Password,
		Net:                  dbConf.Address,
		DBName:               dbConf.Name,
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	dsn := mysqlConf.FormatDSN()
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}

	return db
}
