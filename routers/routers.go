package routers

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"main/api"
	"main/api/club"
	"main/api/user"
	_ "main/docs" //必需
	"main/middleware"
)

func BeginRouters(r *gin.Engine) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Use(middleware.Cors())
	r.POST("/register", club.Register)
	r.GET("/login", club.Login)
	r.POST("callback", api.Callback)
	ClubGroup := r.Group("/club")
	ClubGroup.Use(middleware.JWTAuthMiddleware)
	{
		//发通知或动态
		ClubGroup.POST("post", club.CreatePost)

		ClubGroup.PUT("introduction", club.ModifyClubIntroduction)
		//获取签名
		ClubGroup.GET("signed_url", club.GetSignedUrl)
	}

	UserGroup := r.Group("/user")
	{
		//获取单条通知或动态
		UserGroup.GET("post", user.GetPost)
		UserGroup.GET("club_info", user.ShowClubInfo)
		UserGroup.GET("posts", user.GetPosts)
		UserGroup.GET("posts_from_clubs_user_fellow", user.GetPostsFromClubsUserFollow)
		UserGroup.POST("subscribe", user.Subscribe)
		UserGroup.DELETE("subscribe", user.UnSubscribe)
		UserGroup.POST("collect", user.Collect)
		UserGroup.DELETE("collect", user.CancelCollect)
	}

}
