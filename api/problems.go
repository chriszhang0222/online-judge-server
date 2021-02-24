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

func ProblemDetail(ctx *gin.Context){
	id := ctx.Param("id")
	id_int, _ := strconv.Atoi(id)
	if id_int < 0{
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "Problem does not exist"})
		return
	}
	problem := orm.FindById(id_int)
	if problem == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"msg": "Problem does not exist"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": problem,
	})


}
