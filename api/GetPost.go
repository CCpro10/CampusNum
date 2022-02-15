package api

import (
	"github.com/gin-gonic/gin"
	"time"
)

func UploadPost(c *gin.Context) {
	// @Success 200 {object} json "{"code":200,"data":{models.Post},"msg":"ok"}"
	// @Param is_notice query bool true "是否为通知,是则为true"
}

type ResPost struct {
	//Msg         string `json:"msg"`   //信息 如"获取成功"
	ID          uint      `gorm:"primarykey"json:"id"` //帖子Id
	CreatedAt   time.Time `json:"created_at"`          //创建时间
	UpdatedAt   time.Time `json:"updated_at"`          //更新时间(刚创建时为空)
	Article     string    `json:"article"`             //标题
	Content     string    `json:"content"`             //内容
	PictureAddr []string  `json:"picture_addr"`        //帖子图片的多个可访问地址

	ClubId     uint   `json:"club_id"`                   //社团Id
	ClubName   string `form:"club_name"json:"club_name"` //社团名称
	AvatarAddr string `json:"avatar_addr"`               //头像url地址
}

type ResPosts struct {
	Data []ResPost `json:"data"` //data内包涵多条帖子数据
}

// @Summary 获取单条通知/动态(详情)
// @Accept application/json
// @Produce application/json
// @Param post_id query uint true "帖子的id"
// @Success 200 {object} ResPost
// @Router /user/post [get]
func GetPost(c *gin.Context) {

}

//传入的参数列表
type postParamList struct {
	IsNotice bool  `json:"is_notice"form:"is_notice"`     //"要查询的是否为通知,是则为true,否则为false"
	Page     int64 `json:"page" form:"page" example:"1"`  // 页码
	Size     int64 `json:"size" form:"size" example:"10"` // 每页数据量
}

// @Summary 获取多条通知/动态
// @Accept application/json
// @Produce application/json
// @Param object query postParamList true "参数列表"
// @Success 200 {object} ResPosts "data内有多条post"
// @Router /user/posts [get]
func GetPosts(c *gin.Context) {

}
