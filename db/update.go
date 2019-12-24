package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func AddStudent(student *Student) (interface{}, error) {
	collection := Client.Database(DbName).Collection(ColName)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	res, err := collection.InsertOne(ctx, student)
	return res, err
}

func UpdateStudent(student *StudentUpdateRequest) (interface{}, error) {
	collection := Client.Database(DbName).Collection(ColName)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	filter := bson.M{"email": student.Email}
	update := bson.M{"$set": student}
	res, err := collection.UpdateOne(ctx, filter, update)
	return res, err
}
