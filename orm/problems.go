package orm

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"online-judge-server/global"
	"online-judge-server/model"
)

func GetProblemList(page, pnum int)[]*model.Problem{
	collection := global.MongoClient.Database(global.GlobalConfig.Mongo.Name).Collection("problemmodels")
	options := options.Find()
	skip := pnum * (page - 1)
	options.SetSkip(int64(skip))
	options.SetLimit(int64(pnum))
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
	return lst
}

func InsertProblem(problem *model.Problem) bool{
	collection := global.MongoClient.Database(global.GlobalConfig.Mongo.Name).Collection("problemmodels")
	_, err := collection.InsertOne(context.Background(), bson.M{"name": problem.Name, "desc": problem.Desc, "diff": problem.Diff})
	if err == nil{
		return true
	}
	return false
}

func FindByName(query string)[]*model.Problem{
	collection := global.MongoClient.Database(global.GlobalConfig.Mongo.Name).Collection("problemmodels")
	lst := make([]*model.Problem, 0)
	filter := bson.M{"status": bson.M{"$regex":".*"+query+".*"}}
	data, _ := collection.Find(context.Background(), filter)
	defer data.Close(context.Background())
	for data.Next(context.Background()){
		var res *model.Problem = &model.Problem{}
		err := data.Decode(res)
		if err == nil{
			lst = append(lst, res)
		}
	}
	return lst
}

func FindById(id int)*model.Problem{
	collection := global.MongoClient.Database(global.GlobalConfig.Mongo.Name).Collection("problemmodels")
	filter := bson.M{"id": id}
	result := &model.Problem{}
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return nil
	}
	return result
}

func TotalCount() int64{
	collection := global.MongoClient.Database(global.GlobalConfig.Mongo.Name).Collection("problemmodels")
	count, err := collection.CountDocuments(context.Background(), bson.M{})
	if err != nil{
		return 0
	}
	return count
}
