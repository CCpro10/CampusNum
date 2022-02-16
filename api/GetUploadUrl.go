package api

import (
	"github.com/gin-gonic/gin"
	"main/conf"
	"main/util"
	"net/http"
)

type ResUrl struct {
	Url         string `json:"url"`          //签名url
	CallbackStr string `json:"callback_str"` //回调的字符串
	PictureId   string `json:"picture_id"`   //图片Id
}

type UrlParamList struct {
	PictureName string `json:"picture_name"` //要发送的图片名
	Type        string `json:"type"`         //"post_picture"或"avatar"
}

// @Summary 获取上传图片或头像的签名
// @Description 获取上传图片的url及回调字符串,图片Id
// @Produce json
// @Param Authorization header string false "Bearer 用户令牌 例:Bearer fbhraewifvg43uwerfaewobf"
// @Param object query UrlParamList true "传入参数"
// @Success 200 {object} ResUrl
// @Router /club/signed_url [get]
func GetSignedUrl(c *gin.Context) {

	//获取sts
	credentials := util.GetAssumeRole(
		conf.Config.Oss.RegionId,
		conf.Config.Oss.AccessKeyId,
		conf.Config.Oss.AccessKeySecret,
		conf.Config.Oss.OssUploadRoleArn,
		"XiaoChen").Credentials

	url, callbackStr, _ := util.GetSignedUrl(
		credentials.SecurityToken,
		credentials.AccessKeyId,
		credentials.AccessKeySecret,
		"屏幕截图 2021-11-08 234931.png",
		2,
		"post_picture/")

	c.JSON(http.StatusOK, gin.H{
		"url":          url,
		"callback_str": callbackStr,
	})
}
