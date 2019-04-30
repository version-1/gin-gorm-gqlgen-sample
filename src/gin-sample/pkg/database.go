package pkg

import (
	"gin-sample/configs"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func Connect(env string) *gorm.DB {
	dbConf := configs.GetConfig(env).Database
	mysqlConf := mysql.Config{
		User:                 dbConf.User,
		Passwd:               dbConf.Password,
		Net:                  dbConf.Address,
		DBName:               dbConf.Name,
		AllowNativePasswords: true,
	}

	dsn := mysqlConf.FormatDSN()
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}

	return db
}
