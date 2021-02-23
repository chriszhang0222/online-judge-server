package initialize
import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"online-judge-server/global"
)


func InitConfig() {
	configFileName := "./config.yaml"
	v := viper.New()
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		zap.S().Error("Error when read config yaml", err.Error())
		panic(err)
	}
	if err := v.Unmarshal(global.GlobalConfig); err != nil {
		panic(err)
	}
}
