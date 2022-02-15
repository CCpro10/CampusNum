package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Conf struct {
	Redis struct {
		Addr     string `yaml:"addr"`
		Password string `yaml:"password"`
		DB       int    `yaml:"db"`
	}
	MYSQL struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Addr     string `yaml:"addr"`
		Database string `yaml:"database"`
	}
	Jwt struct {
		TokenExpireDuration int    `yaml:"tokenexpireduration"` //小时为单位
		Secret              string `yaml:"secret"`
	}
	Deploy struct {
		Secret   int64  `yaml:"secret"`
		LocalIp  string `yaml:"localip"`
		ServerIp string `yaml:"serverip"`
	}
	Oss struct {
		AccessKeyId       string `yaml:"accesskeyid"`
		AccessKeySecret   string `yaml:"accesskeysecret"`
		EndPoint          string `yaml:"endpoint"`
		OssUploadRoleArn  string `yaml:"ossuploadrolearn"`
		FullAccessRoleArn string `yaml:"fullaccessrolearn"`
		RegionId          string `yaml:"regionid"`
		BucketName        string `yaml:"bucketname"`
	}
}

//获取配置
func GetConf() *Conf {
	var c = Conf{}
	yamlFile, err := ioutil.ReadFile("./conf/conf.yaml")
	if err != nil {
		log.Println(err.Error())
	}

	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Println(err.Error())
	}
	return &c
}

//获取配置文件
var Config = GetConf()
