package routes

import (
	"AAL_time/api"
	"AAL_time/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine{
	ginServer := gin.Default()
	store := cookie.NewStore([]byte("something-very-secret"))
	ginServer.Use(sessions.Sessions("mysession",store))
	ginServer.Use(middleware.Cors())
	v1 := ginServer.Group("start")
	{
		v1.POST("user/register",api.UserRegister)
		v1.POST("user/login",api.UserLogin)

		authed := v1.Group("timeconsumption")
		authed.Use(middleware.JWT())
		{
			authed.POST("create",api.CreateTimeConsumption)
			authed.GET("show/:id",api.ShowTimeConsumption)
			authed.PUT("update/:id",api.UpdateTimeConsumption)
			authed.DELETE("delete/:id",api.DeleteTimeConsumption)
		}

		tags := v1.Group("category")
		tags.Use(middleware.JWT())
		{
			tags.POST("create",api.CreateCategory)
			tags.GET("all",api.ShowAllCategory)
			tags.PUT("update/:id",api.UpdateCategory)
		}

		time := v1.Group("counttime")
		time.Use(middleware.JWT())
		{
			// 统计消耗时间，以周、月、年，或者自定义时间为单位，按照不同标签返回
			time.GET("week",api.Week)
			time.GET("month",api.Month)
			time.GET("year",api.Year)
			time.POST("autotime",api.SelectTime)
		}
	}
	return ginServer
}