package initialize

import (
	"fmt"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"online-judge-server/global"
)

func InitRedis(){
	redisConfig := global.GlobalConfig.Redis
	var err error
	global.RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.PORT),
		Password: "", // no password set
		DB:       redisConfig.DB,  // use default DB
	})
	_, err = global.RedisClient.Ping().Result()
	if err != nil{
		zap.S().Error("Error when connected to redis" + fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.PORT))
		panic(err)
	}
}