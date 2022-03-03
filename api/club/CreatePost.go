package club

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"main/models"
	"net/http"
)

//请求要增加的帖子
type ReqPost struct {
	IsNotice   bool   `form:"is_notice"json:"is_notice"`                       //是否为通知
	Article    string `form:"article"json:"article"validate:"required,min=2"`  //标题,min=2
	Content    string `form:"content"json:"content"validate:"required,min=10"` //内容,min=10
	PictureIds []uint `form:"picture_ids"json:"picture_ids"validate:"max=9"`   //包涵要上传的帖子图片的id的数组,最多9张图
}

//增加帖子的返回结构体
type CreateResponse struct {
	Msg string `json:"msg"` //返回的信息
}

// @Summary 创建活动或动态
// @Produce json
// @Param Authorization header string true "Bearer 用户令牌 例:Bearer fbhraewifvg43uwerfaewobf"
// @Param object formData ReqPost true "请求所需的参数"
// @Success 200 {object} CreateResponse
// @Router /club/post [post]
func CreatePost(c *gin.Context) {
	//绑定参数并检验
	var reqPost ReqPost
	if err := c.ShouldBind(&reqPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "参数绑定失败"})
		return
	}
	validate := validator.New()
	err := validate.Struct(reqPost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "参数不符合规范,err=" + err.Error()})
		return
	}

	//获取用户信息
	ClubId, ok := c.Get("ClubId")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "获取用户信息失败"})
		return
	}

	clubInfo := models.GetClubInfoById(ClubId.(uint))
	//检查图片能否上传
	for _, v := range reqPost.PictureIds {
		if ok := models.CheckPostPictureById(ClubId.(uint), v); !ok {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "图片Id有误"})
			return
		}
	}
	p := models.Post{
		IsNotice: reqPost.IsNotice,
		ClubId:   clubInfo.ID,
		Article:  reqPost.Article,
		Content:  reqPost.Content,
		ClubName: clubInfo.ClubName,
	}
	e := p.CreatePost(clubInfo.ID, reqPost.PictureIds)
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "创建失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "创建成功"})
	return
}
