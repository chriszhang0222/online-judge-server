package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ProblemList(ctx *gin.Context){
	res := gin.H{
		"msg": "OK",
	}
	ctx.JSON(http.StatusOK, res)
}
