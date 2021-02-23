package main
import (
	"fmt"
	"online-judge-server/global"
	"online-judge-server/initialize"
)
func main(){
	initialize.InitConfig()
	fmt.Println(global.GlobalConfig.Mongo)
}
