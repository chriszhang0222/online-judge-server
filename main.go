package main

import (
	"flag"
	"fmt"
	"go.uber.org/zap"
	"online-judge-server/initialize"
	"os"
	"os/signal"
	"syscall"
)
var port int = 8888

func parsePort(){
	flag.IntVar(&port, "port", 8888, "Cache server port")
	flag.Parse()
}
func main(){
	initialize.InitConfig()
	initialize.InitLogger()
	parsePort()
	Router := initialize.InitRouter()
	go func() {
		if err := Router.Run(fmt.Sprintf(":%d", port)); err != nil {
			zap.S().Panic("serve error", err.Error())
		}
		zap.S().Debugf("serve online judge server at %d", port)
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit


}
