package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"main/models"
	"main/routers"
)

func Setup() {
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	//连接数据库
	models.InitMySQL()
}

func main() {
	Setup()
	r := gin.Default()
	routers.BeginRouters(r)

	e := r.Run(":8008")
	if e != nil {
		panic(e.Error())
	}
}
