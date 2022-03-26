package main

import (
	"awesomeProject1/OAuth2"
	"awesomeProject1/Posts"
	"awesomeProject1/Redis"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	{
		Redis.InitRedis()
	} //init

	r.GET("/oauth/oauth2/redirect", OAuth2.GetLoginFunc)
	r.POST("/post/update", Posts.GetUpdatePostFunc)
	r.POST("/post/sub", Posts.SubPost)

	r.Run(":80")

}
