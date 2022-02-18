package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"main/models"
	"net/http"
	"strconv"
	"strings"
)

//用户登录信息
type ClubLogin struct {
	ClubNameOrId string `form:"club_name_or_id"json:"club_name_or_id"validate:"required,min=2"` //社团名或社团账号,至少2位
	Password     string `form:"password"json:"password"validate:"required,min=6"`               //密码,至少6位
}

type ResLogin struct {
	Msg   string `json:"msg"`   //信息
	Token string `json:"token"` //token
}

// @Summary 登录社团账号
// @Produce json
// @Param object formData ClubLogin true "登录所需要的参数"
// @Success 200 {object} ResLogin
// @Router /login [post]
func Login(c *gin.Context) {
	var requestUser ClubLogin
	err := c.ShouldBind(&requestUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "参数绑定失败",
		})
		return
	}

	requestUser.ClubNameOrId = strings.TrimSpace(requestUser.ClubNameOrId)
	requestUser.Password = strings.TrimSpace(requestUser.Password)
	validate := validator.New()        // 创建验证器
	err = validate.Struct(requestUser) // 执行验证
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": "参数不符合规范,err=" + err.Error()})
		return
	}
	//判断是否为账号
	_, err = strconv.ParseInt(requestUser.ClubNameOrId, 10, 64)
	if err != nil { //输入的是社团名称
		if !models.ExistClub("club_name", requestUser.ClubNameOrId) {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "此名称不存在"})
			return
		}
		if Token, ok := models.VerifyPassword("club_name", requestUser.ClubNameOrId, requestUser.Password); !ok {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "密码输入错误"})
			return
		} else {
			c.JSON(http.StatusOK, ResLogin{
				"登录成功",
				Token,
			})
			return
		}
	}
	//输入的是账号
	if !models.ExistClub("club_id", requestUser.ClubNameOrId) {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "此账号不存在"})
		return
	}
	if Token, ok := models.VerifyPassword("club_id", requestUser.ClubNameOrId, requestUser.Password); !ok {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "密码输入错误"})
		return
	} else {
		c.JSON(http.StatusOK, ResLogin{
			"登录成功",
			Token,
		})
		return
	}
}
