package routers

import (
	"github.com/gin-gonic/gin"
	"online-judge-server/api"
)

func InitProblemsRouter(group *gin.RouterGroup){
	GoodsRouter := group.Group("problems")
	GoodsRouter.GET("", api.ProblemList)

}
