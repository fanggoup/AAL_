package api

import (
	"AAL_time/package/utils"
	"AAL_time/service"

	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context){
	var userRegister service.UserService
	if err := c.ShouldBind(&userRegister);err == nil{
		res := userRegister.Register()
		c.JSON(200,res)
	}else{
		c.JSON(400,err)
		utils.LogrusObj.Info("用户绑定失败",err)
	}
}

func UserLogin(c *gin.Context){
	var userLogin service.UserService
	if err := c.ShouldBind(&userLogin);err == nil{
		res := userLogin.Login()
		c.JSON(200,res)
	}else{
		c.JSON(400,err)
		utils.LogrusObj.Info("用户登录",err)
	}
}