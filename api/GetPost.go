package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"main/models"
	"net/http"
	"strconv"
	"time"
)

type ResponsePost struct {
	//Msg         string `json:"msg"`   //信息 如"获取成功"
	ID          uint      `gorm:"primarykey"json:"id"` //帖子Id
	CreatedAt   time.Time `json:"created_at"`          //创建时间
	UpdatedAt   time.Time `json:"updated_at"`          //更新时间(刚创建时为空)
	Article     string    `json:"article"`             //标题
	Content     string    `json:"content"`             //内容
	PictureAddr []string  `json:"picture_addr"`        //帖子图片的多个可访问地址
	IsNotice    bool      `json:"is_notice"`           //是否为通知

	ClubId     uint   `json:"club_id"`                   //社团Id
	ClubName   string `form:"club_name"json:"club_name"` //社团名称
	AvatarAddr string `json:"avatar_addr"`               //头像url地址
}

type ResponsePosts struct {
	Data []ResponsePost `json:"data"` //data内包涵多条帖子数据
}

// @Summary 获取单条通知/动态(详情)
// @Produce json
// @Param post_id query uint true "帖子的id,min=1"
// @Success 200 {object} ResponsePost
// @Router /user/post [get]
func GetPost(c *gin.Context) {
	//获取参数,检查格式
	req := c.Query("post_id")
	reqPostId, e := strconv.ParseUint(req, 10, 32)
	if e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "post_id参数格式错误"})
		return
	}
	post, ok := models.GetPostById(reqPostId)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "post_id不存在"})
		return
	}
	//把post的值放到rep中
	bytes, _ := json.Marshal(post)
	var rsp ResponsePost
	_ = json.Unmarshal(bytes, &rsp)
	//获取头像
	rsp.AvatarAddr, _ = models.GetAvatarAddrByClubId(post.ClubId)
	//获取图片
	rsp.PictureAddr, _ = models.GetPictureAddrByPostId(post.ID)

	c.JSON(http.StatusOK, rsp)
	return

}

//传入的参数列表
type postParamList struct {
	IsNotice bool  `json:"is_notice"form:"is_notice"`     //"要查询的是否为通知,是则为true,否则为false"
	Page     int64 `json:"page" form:"page" example:"1"`  // 页码,最小为1
	Size     int64 `json:"size" form:"size" example:"10"` // 每页数据量
}

// @Summary 获取最新的多条通知/动态
// @Produce json
// @Param object query postParamList true "参数列表"
// @Success 200 {object} ResponsePosts "data内有多条post"
// @Router /user/posts [get]
func GetPosts(c *gin.Context) {

}
