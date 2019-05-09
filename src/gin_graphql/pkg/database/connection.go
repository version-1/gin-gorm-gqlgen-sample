package database

import (
	"gin_graphql/pkg/config"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"os"
	"sync"
)

var instance *gorm.DB
var once sync.Once

func Connect(env string) *gorm.DB {
	dbConf := config.GetConfig(env).Database
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

func GetInstance() *gorm.DB {
	env := os.Getenv("ENV")
	once.Do(func() {
		instance = Connect(env)
		if env != "production" {
			instance.LogMode(true)
		}
	})
	return instance
}

func Close() {
	if instance == nil {
		return
	}
	instance.Close()
}
