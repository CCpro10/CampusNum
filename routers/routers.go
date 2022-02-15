package routers

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"main/api"
	_ "main/docs" //必需
	"main/middleware"
)

func BeginRouters(r *gin.Engine) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Use(middleware.Cors())
	r.POST("/register", api.Register)
	r.POST("/login", api.Login)
	ClubGroup := r.Group("/club")
	ClubGroup.Use(middleware.JWTAuthMiddleware)
	{
		//发通知或动态
		ClubGroup.POST("post", api.UploadPost)

		ClubGroup.GET("signed_url", api.GetSignedUrl)
	}

	UserGroup := r.Group("/user")
	{
		//获取单条通知或动态
		UserGroup.GET("post", api.GetPost)

		UserGroup.GET("posts", api.GetPosts)
	}

}
