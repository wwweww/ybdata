package service

import (
	"YBData/config"
	"YBData/module"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	host := config.DBCfg.GetHost()
	port := config.DBCfg.GetPort()
	database := config.DBCfg.GetDatabase()
	username := config.DBCfg.GetUsername()
	password := config.DBCfg.GetPassword()
	charset := config.DBCfg.GetCharset()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		database,
		charset,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database, err: " + err.Error())
	}
	_ = db.AutoMigrate(module.StData{})

	DB = db
}
func GetDB() *gorm.DB {
	return DB
}
