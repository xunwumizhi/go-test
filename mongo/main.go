package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
)

var once sync.Once
var ctx context.Context
var cancel context.CancelFunc
var client *mongo.Client

func init() {
	once.Do(func() {
		ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
		var err error
		client, err = mongo.Connect(ctx, clientOptions)
		if err != nil {
			log.Fatalln(err)
		}
	})
}

func main() {
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatalln(err)
		}

		cancel()
	}()

	if err := client.Ping(context.TODO(), nil); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("connection is ok")

	func() {
		// if dbs, err := client.ListDatabaseNames(context.TODO(), bson.D{{Key: "empty", Value: false}}); err != nil {
		// 	log.Fatalln(err)
		// } else {
		// 	fmt.Println("db: ", dbs)
		// }

		// insert()
		query()
		queryByAll()
		// update()
	}()
}

type Info struct {
	Name string
	Age  int
}

func insert() {
	coll := client.Database("user").Collection("info")
	doc := Info{"Tom", 3}
	insertResult, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(insertResult)

	docs := []interface{}{Info{"Jerry", 2}, Info{"Bob", 4}}
	insertManyRes, err := coll.InsertMany(context.TODO(), docs)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(insertManyRes)

}

func query() {
	coll := client.Database("user").Collection("info")
	cur, err := coll.Find(context.TODO(), bson.D{{"age", 1}})
	if err != nil {
		log.Fatalln(err)
	}
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		fmt.Println("RemainingBatchLength: ", cur.RemainingBatchLength())
		e := &Info{}
		cur.Decode(e)
		fmt.Printf("query row: %+v\n", *e)
	}
	if err := cur.Err(); err != nil {
		log.Fatalln("cursor error: ", err)
	}
}

func queryByAll() {
	coll := client.Database("user").Collection("info")
	cur, err := coll.Find(context.TODO(), bson.D{{"age", 1}})
	if err != nil {
		log.Fatalln(err)
	}
	defer cur.Close(context.TODO())

	fmt.Println("RemainingBatchLength: ", cur.RemainingBatchLength())
	fmt.Println("next: ", cur.Next(context.TODO()))
	res := &[]*Info{}
	if err := cur.All(context.TODO(), res); err != nil {
		log.Fatalln(err)
	}
	bs, _ := json.Marshal(res)
	fmt.Println("cursor.All: ", string(bs))
	fmt.Println("RemainingBatchLength: ", cur.RemainingBatchLength())
	fmt.Println("next: ", cur.Next(context.TODO()))

}

func update() {
	filter := bson.D{
		{"name", "sample"},
	}
	update := bson.D{
		{"$set", bson.D{{"email", "sample@email.com"}}},
	}
	opts := options.Update().SetUpsert(true)

	coll := client.Database("user").Collection("info")
	res, err := coll.UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%+v", *res)
}
