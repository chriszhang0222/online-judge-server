package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"online-judge-server/orm"
	"strconv"
)

func ProblemList(ctx *gin.Context){
	pages := ctx.DefaultQuery("page", "1")
	pages_int, _ := strconv.Atoi(pages)
	perNums := ctx.DefaultQuery("pnum", "10")
	perNumsInt, _ := strconv.Atoi(perNums)
	problem_list := orm.GetProblemList(pages_int, perNumsInt)


	res := gin.H{
		"msg": "OK",
		"data": problem_list,
	}
	ctx.JSON(http.StatusOK, res)
}
