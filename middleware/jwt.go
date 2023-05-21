package middleware

import (
	"AAL_time/package/e"
	"AAL_time/package/utils"
	"time"

	"github.com/gin-gonic/gin"
)

// JWT token验证中间件
func JWT() gin.HandlerFunc{
	return func(c *gin.Context){
		var code = 200
		token := c.GetHeader("Authorization")
		if token == ""{
			code = 404
		}else{
			claims,err := utils.ParseToken(token)
			if err != nil {
				code = e.ErrorAuthCheckTokenFail
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ErrorAuthCheckTokenTimeout
			}
		}
		if code != e.SUCCESS{
			c.JSON(400,gin.H{
				"status":code,
				"msg":e.GetMsg(code),
				"data":"身份信息已过期，请重新登录",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}