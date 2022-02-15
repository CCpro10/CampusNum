package util

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
)

//获取扮演的角色的STS临时凭证
func GetAssumeRole(RegionId string, AccessKeyId string, AccessKeySecret string, RoleArn string, RoleSessionName string) *sts.AssumeRoleResponse {

	//构建一个阿里云客户端, 用于发起请求。
	//构建阿里云客户端时，需要设置AccessKey ID和AccessKey Secret。
	client, err := sts.NewClientWithAccessKey(RegionId, AccessKeyId, AccessKeySecret)

	//构建请求对象。
	request := sts.CreateAssumeRoleRequest()
	request.Scheme = "https"

	//设置参数。关于参数含义和设置方法，请参见《API参考》。
	request.RoleArn = RoleArn
	request.RoleSessionName = RoleSessionName

	//发起请求，并得到响应。
	response, err := client.AssumeRole(request)
	if err != nil {
		fmt.Print(err.Error())
	}

	return response
}
