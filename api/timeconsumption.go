package api

import (
	"AAL_time/service"

	"AAL_time/package/utils"

	"github.com/gin-gonic/gin"
)

// 创建时间消耗表
func CreateTimeConsumption(c *gin.Context) {
	createService := service.CreateTimeConsumption{}
	chaim,_ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&createService);err == nil{
		res := createService.Create(chaim.Id)
		c.JSON(200,res)
	}else{
		c.JSON(400,ErrorResponse(err))
		utils.LogrusObj.Info(err)
	}
}

// 展示时间消耗表
func ShowTimeConsumption(c *gin.Context){
	showService := service.ShowTimeConsumption{}
	if err := c.ShouldBind(&showService);err == nil{
		res := showService.Show(c.Param("id"))
		c.JSON(200,res)
	}else{
		c.JSON(400,ErrorResponse(err))
		utils.LogrusObj.Info(err)
	}
}

// 修改时间消耗表
func UpdateTimeConsumption(c *gin.Context){
	updateService := service.UpdateTimeConsumption{}
	if err := c.ShouldBind(&updateService);err == nil{
		res := updateService.Update(c.Param("id"))
		c.JSON(200,res)
	}else{
		c.JSON(400,ErrorResponse(err))
		utils.LogrusObj.Info(err)
	}
}

// 删除时间消耗表
func DeleteTimeConsumption(c *gin.Context){
	deleteSerice := service.DeleteTimeConsumption{}
	if err := c.ShouldBind(&deleteSerice);err == nil{
		res := deleteSerice.Delete(c.Param("id"))
		c.JSON(200,res)
	}else{
		c.JSON(400,ErrorResponse(err))
		utils.LogrusObj.Info(err)
	}
}



