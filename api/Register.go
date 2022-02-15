package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main/conf"
	"net/http"
	"time"
)

//用户注册信息
type ClubRegister struct {
	InvitationCode int64  `form:"invitation_code"`           //邀请码
	ClubId         int64  `form:"club_id"`                   //社团登录账号
	ClubName       string `form:"club_name"json:"club_name"` //社团名称
	Password       string `form:"password"json:"password"`   //密码
	Password2      string `form:"password2"json:"password2"` //密码

}

//判断邀请码是否有效
func verifyCode(code int64) bool {
	m := int(time.Now().Month())
	d := time.Now().Day()
	c := (int64(m*d + 7)) * conf.Config.Deploy.Secret
	fmt.Println(c)
	if c != code {
		return false
	}
	return true
}

//注册路由
func Register(c *gin.Context) {
	//从请求中把数据取出
	var requestUser ClubRegister
	err := c.ShouldBind(&requestUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "注册参数绑定失败:" + err.Error()})
		return
	}
	if verifyCode(requestUser.InvitationCode) {
		c.JSON(http.StatusBadRequest, gin.H{"msg": ":邀请码无效或过期"})
		return
	}
	//
	////验证学号,邮件,密码的结构
	//if len(requestUser.StudentId) != 10 {
	//	c.JSON(http.StatusBadRequest, gin.H{"msg": "学号格式错误"})
	//	return
	//}
	//if util.CheckEmail(requestUser.Email) == false {
	//	c.JSON(http.StatusBadRequest, gin.H{"msg": "邮箱格式不正确,请确保输入有效的邮箱"})
	//	return
	//}
	//if len(requestUser.Password) < 6 {
	//	c.JSON(http.StatusBadRequest, gin.H{"msg": "密码长度不能小于六位"})
	//	return
	//}
	//
	////判断在数据库中邮箱是否存在
	//var user model.UserInfo
	////查询并赋值给user
	//dao.DB.Where("email=?", requestUser.Email).First(&user)
	//if user.ID != 0 {
	//	c.JSON(http.StatusBadRequest, gin.H{"msg": "此邮箱已被注册"})
	//	return
	//}
	//
	////查询Redis,检测验证码是否正确,是否过期
	//value, err := dao.RedisDB.Get(dao.CTX, requestUser.Email).Result()
	//
	////如果查不到此邮箱对应验证码
	//if err == redis.Nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"msg": "验证码错误或已经过期"})
	//	return
	//} else if err != nil { //查询失败了
	//	c.JSON(http.StatusInternalServerError, gin.H{"msg": "服务器查询验证码失败"})
	//	return
	//}
	//
	////取出value里的验证码vcode
	//parts := strings.SplitN(value, " ", 2)
	//vcode := parts[0]
	//
	////通过邮箱查到的验证码和用户输入的验证码不一样
	//if vcode != requestUser.VerifyCode {
	//	c.JSON(http.StatusBadRequest, gin.H{"msg": "验证码错误或已经过期"})
	//	return
	//}
	//
	//// 对密码加密
	//hashPassword, err := bcrypt.GenerateFromPassword([]byte(requestUser.Password), bcrypt.DefaultCost)
	//if err != nil {
	//	c.JSON(http.StatusUnprocessableEntity, gin.H{"msg": "用户密码加密失败"})
	//	return
	//}
	//
	////创建用户实例,存入注册信息
	//var userinfo = model.UserInfo{
	//	Email:     requestUser.Email,
	//	StudentId: requestUser.StudentId,
	//	Password:  string(hashPassword),
	//}
	//err = dao.DB.Create(&userinfo).Error
	//if err != nil {
	//	c.JSON(http.StatusOK, gin.H{"msg": "注册信息存入数据库失败:" + err.Error()})
	//} else {
	//	c.JSON(http.StatusOK, gin.H{
	//		"msg":  "注册成功,请你重新登录",
	//		"date": userinfo,
	//	})
	//}
}
