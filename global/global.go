package global

import (
	"go.mongodb.org/mongo-driver/mongo"
	"online-judge-server/config"
	"github.com/go-redis/redis"
)

var(
	GlobalConfig = &config.MainConfig{}
	MongoClient *mongo.Client
	RedisClient *redis.Client
)
