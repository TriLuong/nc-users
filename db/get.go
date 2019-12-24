package db

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func GetAllStudent() (*[]Student, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	filter := bson.M{} //map[string]interface{}
	cur, err := Client.Database(DbName).Collection(ColName).Find(ctx, filter)
	if err != nil {
		log.Printf("Find error: %v", err)
		return nil, err
	}

	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	var students []Student
	err = cur.All(ctx, &students)
	if err != nil {
		log.Printf("cur all error: %v", err)
		return nil, err
	}

	return &students, nil
}

func SearchStudentSimple(req StudentSearchRequest) (*[]Student, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	var filter bson.M
	bs, err := json.Marshal(req)
	err = json.Unmarshal(bs, &filter)
	if err != nil {
		log.Printf("marshal error: %v", err)
	}

	fmt.Println(filter)

	cur, err := Client.Database(DbName).Collection(ColName).Find(ctx, filter)
	if err != nil {
		log.Printf("Find error: %v", err)
		return nil, err
	}

	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	var students []Student
	err = cur.All(ctx, &students)
	if err != nil {
		log.Printf("cur all error: %v", err)
		return nil, err
	}

	return &students, nil
}
