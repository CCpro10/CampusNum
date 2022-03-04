package club

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"log"
	"main/models"
	"main/util"
	"net/http"
	"strconv"
	"strings"
)

//用户注册信息
type ReqRegister struct {
	InvitationCode int64  `form:"invitation_code"json:"invitation_code"validate:"required"`            //邀请码
	Account        int64  `form:"club_id"json:"club_id"validate:"required,min=99999,max=999999999999"` //社团登录账号7-12位
	ClubName       string `form:"club_name"json:"club_name"validate:"required,min=2"`                  //社团名称
	Password       string `form:"password"json:"password"validate:"required,min=6,max=32"`             //密码6-32位
	Password2      string `form:"password2"json:"password2"validate:"required,eqfield=Password"`       //	确认密码

}

type RspRegister struct {
	Msg      string `json:"msg"`     //信息
	ClubId   int64  `json:"club_id"` //社团Id
	ClubName string `json:"club_name"`
	Password string `json:"password"` //用户密码
}

// @Summary 注册社团账号
// @Produce json
// @Param object formData ReqRegister true "注册所需要的参数"
// @Success 200 {object} RspRegister
// @Router /register [post]
func Register(c *gin.Context) {
	//从请求中把数据取出
	var requestUser ReqRegister
	err := c.ShouldBind(&requestUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "注册参数绑定失败:" + err.Error()})
		return
	}
	if !util.VerifyCode(requestUser.InvitationCode) {
		c.JSON(http.StatusBadRequest, gin.H{"msg": ":邀请码无效或过期"})
		return
	}
	//去空格
	requestUser.ClubName = strings.TrimSpace(requestUser.ClubName)
	requestUser.Password = strings.TrimSpace(requestUser.Password)
	requestUser.Password2 = strings.TrimSpace(requestUser.Password2)

	validate := validator.New()        // 创建验证器
	err = validate.Struct(requestUser) // 执行验证
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": "参数不符合规范,err=" + err.Error()})
		return
	}

	if models.ExistClub("account", strconv.FormatInt(requestUser.Account, 10)) {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "账号已存在"})
		return
	}

	if models.ExistClub("club_name", string(requestUser.ClubName)) {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "社团名称重复"})
		return
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(requestUser.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"msg": "用户密码加密失败"})
		return
	}

	//创建用户实例,存入注册信息
	var clubInfo = models.ClubInfo{
		Account:  requestUser.Account,
		ClubName: requestUser.ClubName,
		Password: string(hashPassword),
	}

	if !models.CreateClubInfo(&clubInfo) {
		c.JSON(http.StatusOK, gin.H{"msg": "注册信息存入数据库失败"})
	}

	c.JSON(http.StatusOK, RspRegister{
		"注册成功,请重新登陆",
		requestUser.Account,
		requestUser.ClubName,
		requestUser.Password,
	})

}
