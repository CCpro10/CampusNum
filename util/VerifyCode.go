package util

import (
	"main/conf"
	"time"
)

//判断邀请码是否有效
func VerifyCode(code int64) bool {
	m := int(time.Now().Month())
	d := time.Now().Day()
	c := (int64(m*d + 7)) * conf.Config.Deploy.Secret

	if c != code {
		return false
	}
	return true
}
