package routers

import (
	"github.com/gin-gonic/gin"
	"online-judge-server/api"
)

func InitProblemsRouter(group *gin.RouterGroup){
	Router := group.Group("problems")
	{
		Router.GET("", api.ProblemList)
		Router.POST("", api.AddProblem)
		Router.GET("/:id", api.ProblemDetail)
	}
	Router2 := group.Group("data")
	{
		Router2.GET("count", api.TotalCount)
		Router2.GET("goroutine", api.SystemStateHandler)
	}


}
