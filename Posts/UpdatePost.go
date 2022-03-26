package Posts

import (
	"awesomeProject1/Redis"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUpdatePostFunc(c *gin.Context) {

	//now update a Post

	postid := c.PostForm("postid")

	users := Redis.LRange(postid+":subscr", 0, -1)

	for _, user := range users {
		Redis.LPush(user+":msg", "UPDATE:"+postid)
	} //Redis -> userName:msg <- updatePostMessage

	c.JSON(http.StatusOK, gin.H{
		"state": 10000,
		"data":  "ok",
	})

}
