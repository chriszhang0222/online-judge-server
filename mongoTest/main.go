package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"online-judge-server/global"
	"online-judge-server/initialize"
	"online-judge-server/model"
)

func main(){
	initialize.InitLogger()
	initialize.InitConfig()
	collection := global.MongoClient.Database(global.GlobalConfig.Mongo.Name).Collection("problemmodels")
	options := options.Find()
	options.SetLimit(1)
	options.SetSkip(0)
	data, _ := collection.Find(context.Background(), bson.D{}, options)
	defer data.Close(context.Background())
	lst := make([]*model.Problem, 0)
	for data.Next(context.Background()){
		var res *model.Problem = &model.Problem{}
		err := data.Decode(res)
		if err == nil{
			lst = append(lst, res)
		}
	}
	fmt.Println(len(lst))


}
