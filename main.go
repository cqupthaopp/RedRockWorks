package main

import (
	"awesomeProject1/OAuth2"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/oauth/oauth2/redirect", OAuth2.GetLoginFunc)

	r.Run(":80")

}
