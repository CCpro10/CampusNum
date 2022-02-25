package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"main/models"
	"net/http"
)

type ResponsePosts struct {
	Data []ResponsePost `json:"data"` //data内包涵多条帖子数据
}

//传入的参数列表
type postParamList struct {
	IsNotice bool `json:"is_notice"form:"is_notice"`                                     //"要查询的是否为通知,是则为true,否则为false"
	Page     int  `json:"page" form:"page" example:"1"validate:"required,min=1"`         // 页码,最小为1
	Size     int  `json:"size" form:"size" example:"10"validate:"required,min=1,max=40"` //每页数据量,最大为40
}

// @Summary 获取最新的多条通知/动态
// @Produce json
// @Param object query postParamList true "参数列表"
// @Success 200 {object} ResponsePosts "data内有多条post"
// @Router /user/posts [get]
func GetPosts(c *gin.Context) {
	var request postParamList
	if e := c.ShouldBind(&request); e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "参数绑定失败"})
		return
	}
	log.Println(request)

	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "参数不符合规范,err=" + err.Error()})
		return
	}

	posts, _ := models.GetPosts(request.IsNotice, request.Page, request.Size)

	c.JSON(http.StatusOK, posts)
	return

}
