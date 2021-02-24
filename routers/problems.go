package routers

import (
	"github.com/gin-gonic/gin"
	"online-judge-server/api"
)

func InitProblemsRouter(group *gin.RouterGroup){
	Router := group.Group("problems")
	Router.GET("", api.ProblemList)
	Router.GET("/:id", api.ProblemDetail)

}
