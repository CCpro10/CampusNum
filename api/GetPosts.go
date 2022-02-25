package api

import (
	"encoding/json"
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

	var resPosts []ResponsePost
	for _, post := range *posts {
		resPosts = append(resPosts, *BindPost(&post))
	}

	c.JSON(http.StatusOK, resPosts)
	return

}

//绑定帖子的图片已经社团头像
func BindPost(post *models.Post) (rspPost *ResponsePost) {
	//把post的值放到rep中
	bytes, _ := json.Marshal(post)
	_ = json.Unmarshal(bytes, &rspPost)
	//获取头像
	rspPost.AvatarAddr, _ = models.GetAvatarAddrByClubId(post.ClubId)
	//获取图片
	rspPost.PictureAddr, _ = models.GetPictureAddrByPostId(post.ID)
	return
}
