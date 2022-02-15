package models

import (
	"main/models/initial"
	"time"
)

//社团信息
type ClubInfo struct {
	ID           uint  `gorm:"primarykey"`
	ClubId       int64 //社团注册时的账号
	CreatedAt    time.Time
	AvatarAddr   string `json:"avatar_addr"`                     //社团头像url地址
	Introduction string `form:"introduction"json:"introduction"` //社团简介
	Email        string `form:"email"json:"email"`               //电子邮箱
	ClubName     string `form:"club_name"json:"club_name"`       //社团名称
	Password     string `form:"password"json:"password"`         //密码

}

//注册时创建社团账号
func CreateClubInfo(data map[string]interface{}) bool {

	//创建时的具体类型一定要明确
	initial.DB.Create(&ClubInfo{
		ClubId:   data["club_id"].(int64),
		ClubName: data["club_name"].(string),
		Password: data["password"].(string),
	})

	return true
}
