package main

import "github.com/gin-gonic/gin"

type User struct {
	Uid int `json:"uid" form:"uid"` //用户id
}

func main() {
	r := gin.Default()
	r.GET("/bind/:Uid", func(c *gin.Context) {
		var u User
		e := c.BindUri(&u)
		//e := c.ShouldBind(&u)
		if e == nil {
			c.JSON(200, u)
		}
	})
	_ = r.Run()
}
