package api

import (
	"AAL_time/package/utils"
	"AAL_time/serializer"
	"AAL_time/service"
	// "fmt"

	"fmt"

	"github.com/gin-gonic/gin"
)

func Week(c *gin.Context){
	// 获取用户所有标签
	showallService := service.ShowAllCategory{}
	chaim,_ := utils.ParseToken(c.GetHeader("Authorization"))
	var res serializer.Response
	
	if err := c.ShouldBind(&showallService);err == nil{
		res = showallService.ShowAll(chaim.Id)
	}

	weekStart,now := utils.WeekStart()
	countTime := service.Start(res,weekStart.Unix(),now.Unix())
	c.JSON(200,countTime)
}

func Month(c *gin.Context){
	// 获取用户所有标签
	showallService := service.ShowAllCategory{}
	chaim,_ := utils.ParseToken(c.GetHeader("Authorization"))
	var res serializer.Response
	if err := c.ShouldBind(&showallService);err == nil{
		res = showallService.ShowAll(chaim.Id)
	}

	monthStart,now := utils.MonthStart()
	countTime := service.Start(res,monthStart.Unix(),now.Unix())
	c.JSON(200,countTime)
}


func Year(c *gin.Context){
	// 获取用户所有标签
	showallService := service.ShowAllCategory{}
	chaim,_ := utils.ParseToken(c.GetHeader("Authorization"))
	var res serializer.Response
	if err := c.ShouldBind(&showallService);err == nil{
		res = showallService.ShowAll(chaim.Id)
	}

	yearStart,now := utils.YearStart()
	countTime := service.Start(res,yearStart.Unix(),now.Unix())
	c.JSON(200,countTime)
}

func SelectTime(c *gin.Context){
	selectService := service.SelectTime{}
	chaim,_ := utils.ParseToken(c.GetHeader("Authorization"))
	var res serializer.Response
	err := c.ShouldBind(&selectService)
	if err == nil{
		res = selectService.ShowAll(chaim.Id)
	}
	fmt.Println(selectService.StartTime,selectService.EndTime)
	countTime := service.Start(res,selectService.StartTime,selectService.EndTime)
	c.JSON(200,countTime)
}