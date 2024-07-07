package database

import (
	"fmt"
	"log"
	"taobao_backend/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	instance *gorm.DB
)

func InitMysql() {
	conf := config.Cfg.Mysql
	// 构造数据源名称 (DSN)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.User,
		conf.Pass,
		conf.Address,
		conf.Port,
		conf.DbName,
	)
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}
	if db, err := db.DB(); err != nil {
		log.Panic(err)
	} else {
		err = db.Ping()
		if err != nil {
			log.Panic(err)
		}
		log.Println("connect to MySQL success")
	}
	instance = db.Debug()

	_ = instance.AutoMigrate(users{})
	_ = instance.AutoMigrate(cloths{})
}
