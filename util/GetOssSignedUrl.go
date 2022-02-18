package util

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"log"
	"main/conf"
	"path"
	"time"
)

//通过临时角色，生成上传图片的临时URL,和回调的callbackStr
func GetSignedUrl(SecurityToken string, AccessKeyId string, AccessKeySecret string, objectName string, PictureId uint, objectPath string) (signedUrl string, callbackStr string, e error) {
	yourEndpoint := "oss-cn-shenzhen.aliyuncs.com"
	client, e := oss.New(yourEndpoint, AccessKeyId, AccessKeySecret, oss.SecurityToken(SecurityToken))
	if e != nil {
		return
	}
	// 填写Bucket名称
	bucketName := conf.Config.Oss.BucketName

	// 获取存储空间。
	bucket, e := client.Bucket(bucketName)
	if e != nil {
		return
	}
	suffix := path.Ext(objectName)                             //后缀
	t := time.Now().UnixNano()                                 //随机名
	pictureName := objectPath + fmt.Sprintf("%d%s", t, suffix) //图片名
	callbackStr, e = callback(PictureId, pictureName)
	if e != nil {
		return
	}
	signedUrl, e = bucket.SignURL(pictureName, oss.HTTPPut, 3600, []oss.Option{
		oss.Callback(callbackStr),
	}...)
	return signedUrl, callbackStr, e
}

func callback(id uint, pictureName string) (string, error) {
	var callBackStruct = struct {
		CallbackUrl      string `json:"callbackUrl"`
		CallbackBody     string `json:"callbackBody"`
		CallbackBodyType string `json:"callbackBodyType"`
	}{
		"http://" + conf.Config.Deploy.ServerIp + "/callback",
		"size=${size}" + "&picture_id=" + fmt.Sprint(id) + "&picture_name=" + pictureName,
		"application/x-www-form-urlencoded",
	}
	log.Println(callBackStruct)
	callbackStr, e := json.Marshal(callBackStruct)
	if e != nil {
		return "", e
	}
	return base64.StdEncoding.EncodeToString(callbackStr), e
}

//func (a *async) SignUrl(id uint, suffix string) (string, uint, string, error) {
//	callbackStr, num, e := a.callback(id)
//	if e != nil {
//		return "", 0, "", e
//	}
//	key, e := b.SignURL(a.path+fmt.Sprintf("%d.%s", id, suffix), oss.HTTPPut, 300, []oss.Option{
//		oss.Callback(callbackStr),
//	}...)
//	return key, num, callbackStr, nil
//}

//修改为绑定的域名,获取可预览的地址
func GetPreviewableUrl(oldUrl string) (url string) {
	url = oldUrl[:len("http://incu-campus-num")] +
		".ncuos.com" +
		oldUrl[len("http://incu-campus-num.oss-cn-shenzhen.aliyuncs.com"):]
	return

}
