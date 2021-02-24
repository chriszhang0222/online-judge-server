package initialize

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	mongoAddr := fmt.Sprintf("mongodb://%s:%d", global.GlobalConfig.Mongo.Host, global.GlobalConfig.Mongo.PORT)
	clientOptions := options.Client().ApplyURI(mongoAddr)
	var err error
	global.MongoClient, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil{
		zap.S().Error(err)
		return
	}
	err = global.MongoClient.Ping(context.Background(), nil)
	if err != nil {
		zap.S().Error(err)
		return
	}
	zap.S().Info("Connect to " + mongoAddr)
}
