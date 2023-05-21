package api

import (
	"AAL_time/package/utils"
	"AAL_time/service"

	"github.com/gin-gonic/gin"
)

func CreateCategory(c *gin.Context){
	createService := service.CreateCategory{}
	chaim,_ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&createService);err == nil{
		res := createService.Create(chaim.Id)
		c.JSON(200,res)
	}else{
		c.JSON(400,ErrorResponse(err))
		utils.LogrusObj.Info(err)
	}
}

// 用户所有的标签
func ShowAllCategory(c *gin.Context){
	showallService := service.ShowAllCategory{}
	chaim,_ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&showallService);err == nil{
		res := showallService.ShowAll(chaim.Id)
		c.JSON(200,res)
	}else{
		c.JSON(400,ErrorResponse(err))
		utils.LogrusObj.Info(err)
	}
}

func UpdateCategory(c *gin.Context){
	updateService := service.UpdateCategory{}
	if err := c.ShouldBind(&updateService);err == nil{
		res := updateService.Update(c.Param("id"))
		c.JSON(200,res)
	}else{
		c.JSON(400,ErrorResponse(err))
		utils.LogrusObj.Info(err)
	}

}


