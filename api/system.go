package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime"
)

func SystemStateHandler(ctx *gin.Context){
	numGoroutine := runtime.NumGoroutine()
	ctx.JSON(http.StatusOK, gin.H{
		"numGoroutine": numGoroutine,
	})
}
