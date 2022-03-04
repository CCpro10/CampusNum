package user

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"main/models"
	"net/http"
	"strconv"
)

type ReqSubscribe struct {
	StudentId string `json:"student_id"form:"student_id"validate:"required,len=10|len=11"` //学号
	ClubId    uint   `json:"club_id"from:"club_id"validate:"required"`                     //社团Id
}

type ResponseSubscribe struct {
	Msg string //信息
}

// @Summary 关注社团
// @Produce json
// @Param object query ReqSubscribe true "社团id"
// @Success 200 {object} ResponseSubscribe
// @Router /user/subscribe [post]
func Subscribe(c *gin.Context) {
	//参数绑定
	var request ReqSubscribe
	clubId, err := strconv.ParseUint(c.Query("club_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "参数绑定失败"})
		return
	}
	request.ClubId = uint(clubId)

	request.StudentId = c.Query("student_id")

	validate := validator.New()
	err = validate.Struct(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "参数不符合规范,err=" + err.Error()})
		return
	}
	//检查社团ID是否存在
	if !models.ExistClub("id", request.ClubId) {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "无效的社团Id"})
		return
	}
	//获取userId
	userId, _ := models.MustGetUserIdByStudentId(request.StudentId)

	//检查有没有关注
	if models.IsSubscribe(request.ClubId, userId) {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "已关注此社团"})
		return
	}

	//建立关注
	err = models.CreateSubscription(request.ClubId, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "关注失败" + err.Error()})
		return
	}

	c.JSON(http.StatusOK, ResponseSubscribe{
		Msg: "关注成功",
	})
	return

}

// @Summary 取消关注社团
// @Produce json
// @Param object query ReqSubscribe true "社团id"
// @Success 200 {object} ResponseSubscribe
// @Router /user/subscribe [delete]
func UnSubscribe(c *gin.Context) {
	//参数绑定
	var request ReqSubscribe
	clubId, err := strconv.ParseUint(c.Query("club_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "参数绑定失败"})
		return
	}
	request.ClubId = uint(clubId)

	request.StudentId = c.Query("student_id")

	validate := validator.New()
	err = validate.Struct(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "参数不符合规范,err=" + err.Error()})
		return
	}
	//检查社团ID是否存在
	if !models.ExistClub("id", request.ClubId) {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "无效的社团Id"})
		return
	}
	//获取userId
	userId, _ := models.GetUserIdByStudentId(request.StudentId)

	//检查有没有关注
	if !models.IsSubscribe(request.ClubId, userId) {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "用户还未关注此社团"})
		return
	}

	//取消关注
	err = models.CancelSubscribe(request.ClubId, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "取消关注失败" + err.Error()})
		return
	}

	c.JSON(http.StatusOK, ResponseSubscribe{
		Msg: "取消关注成功",
	})
	return
}
