package Posts

import (
	"awesomeProject1/Redis"
	"github.com/gin-gonic/gin"
)

func SubPost(c *gin.Context) {

	postID := c.PostForm("postid")
	userName := c.PostForm("username")

	// key = postID:subscr

	Redis.LPush(postID+":subscr", userName) //SubAPost

	return

}
