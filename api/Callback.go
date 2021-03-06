package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"main/models"
	"net/http"
	"strings"
)

type PicCallback struct {
	ClubId      uint   `json:"club_id"form:"club_id"validate:"required"`
	Size        uint64 `json:"size" form:"size" validate:"required"`
	PictureName string `json:"picture_name" form:"picture_name" validate:"required"`
}

func GetUrlByName(name string) string {
	return "http://incu-campus-num.ncuos.com/" + name
}

//处理oss的回调内容,创建临时图片
func Callback(c *gin.Context) {
	var picCallback PicCallback
	if e := c.ShouldBind(&picCallback); e != nil {
		log.Println(e)
		return
	}
	//type
	t := strings.Split(picCallback.PictureName, "/")[0]

	if t == "avatar" {
		var avatar models.Avatar
		avatar.PictureName = picCallback.PictureName
		avatar.PictureAddr = GetUrlByName(avatar.PictureName)
		id, _ := avatar.CreatePicture()
		c.JSON(http.StatusOK, gin.H{"200": "OK", "picture_id": id})
		return
	}
	//创建未和postId绑定的临时图片
	//获取回调信息,创建临时图片
	var picture models.PostPicture
	picture.ClubId = picCallback.ClubId
	picture.PictureName = picCallback.PictureName
	picture.PictureAddr = GetUrlByName(picture.PictureName)
	id, _ := picture.CreatePicture()

	//给前端返回图片Id
	c.JSON(http.StatusOK, gin.H{"200": "OK", "picture_id": id})
	return

}
