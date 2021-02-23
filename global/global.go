package global

import (
	"go.mongodb.org/mongo-driver/mongo"
	"online-judge-server/config"
)

var(
	GlobalConfig = &config.MainConfig{}
	MongoClient *mongo.Client
)
