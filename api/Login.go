package api

import "github.com/gin-gonic/gin"

//用户登录信息
type ClubLogin struct {
	ClubNameOrId string `form:"club_name_or_id"`         //电子邮箱或学号
	Password     string `form:"password"json:"password"` //密码
}

func Login(c *gin.Context) {

}
