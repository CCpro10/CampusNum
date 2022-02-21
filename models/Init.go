package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"main/conf"
)

//mysql,gorm配置
var DB *gorm.DB

func InitMySQL() {

	dsn := conf.Config.MYSQL.Username + ":" +
		conf.Config.MYSQL.Password + "@tcp(" +
		conf.Config.MYSQL.Addr + ")/" +
		conf.Config.MYSQL.Database + "?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}) //这里用短变量声明会有歧义

	if err != nil {
		panic(err)
	}

	//绑定结构体
	if err = DB.AutoMigrate(
		Post{},
		Avatar{},
		ClubInfo{},
		User{},
		PostPicture{},
		Collection{},
		Subscription{},
	); err != nil {
		log.Panicln(err)
	}

}
