package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"main/conf"
	"main/util"
	"net/http"
	"strings"
)

type ResUrl struct {
	Url         string `json:"url"`          //签名url
	CallbackStr string `json:"callback_str"` //回调的字符串
}

type UrlParamList struct {
	PictureName string `json:"picture_name"validate:"required"` //要发送的图片名
	Type        string `json:"type"validate:"required"`         //"post_picture"或"avatar"
}

// @Summary 获取上传图片或头像的签名
// @Description 获取上传图片的url及回调字符串,图片Id
// @Produce json
// @Param Authorization header string false "Bearer 用户令牌 例:Bearer fbhraewifvg43uwerfaewobf"
// @Param object query UrlParamList true "传入参数"
// @Success 200 {object} ResUrl
// @Router /club/signed_url [get]
func GetSignedUrl(c *gin.Context) {
	var request UrlParamList
	request.Type = c.Query("type")
	request.PictureName = c.Query("picture_name")

	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "参数不符合规范,err=" + err.Error()})
		return
	}
	if (request.Type != "post_picture") && (request.Type != "avatar") {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "type参数不符合规范"})
		return
	}
	if len(strings.Split(request.PictureName, ".")) != 2 {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "图片名称不符合规范"})
		return
	}

	//获取用户信息
	ClubId, ok := c.Get("ClubId")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "获取用户信息失败"})
		return
	}
	//获取sts
	credentials := util.GetAssumeRole(
		conf.Config.Oss.RegionId,
		conf.Config.Oss.AccessKeyId,
		conf.Config.Oss.AccessKeySecret,
		conf.Config.Oss.OssUploadRoleArn,
		"XiaoChen").Credentials
	//获取签名和callbackStr
	url, callbackStr, _ := util.ClubId(ClubId.(uint)).GetSignedUrl(
		credentials.SecurityToken,
		credentials.AccessKeyId,
		credentials.AccessKeySecret,
		request.PictureName,
		request.Type+"/")

	c.JSON(http.StatusOK, ResUrl{
		url,
		callbackStr,
	})
}
