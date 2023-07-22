package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strings"
)

// https://dasarpemrogramangolang.novalagung.com/A-mongodb.html
func main() {
	//insert()
	//find()
	//update1()
	//remove()
	//aggregate()
}

func aggregate() {
	db, err := connect1()
	if err != nil {
		log.Fatal(err.Error())
	}

	pipeline := make([]bson.M, 0)
	err = bson.UnmarshalExtJSON([]byte(strings.TrimSpace(`
    [
        { "$group": {
            "_id": null,
            "Total": { "$sum": 1 }
        } },
        { "$project": {
            "Total": 1,
            "_id": 0
        } }
    ]
`)), true, &pipeline)
	if err != nil {
		log.Fatal(err.Error())
	}

	csr, err := db.Collection("student").Aggregate(ctx, pipeline)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer csr.Close(ctx)

	result := make([]bson.M, 0)
	for csr.Next(ctx) {
		var row bson.M
		err := csr.Decode(&row)
		if err != nil {
			log.Fatal(err.Error())
		}

		result = append(result, row)
	}

	if len(result) > 0 {
		fmt.Println("Total :", result[0]["Total"])
	}
}

func remove() {
	db, err := connect1()
	if err != nil {
		log.Fatal(err.Error())
	}

	var selector = bson.M{"name": "John Wick"}
	_, err = db.Collection("student").DeleteOne(ctx, selector)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Remove success!")
}

func update1() {
	db, err := connect1()
	if err != nil {
		log.Fatal(err.Error())
	}

	var selector = bson.M{"name": "Wick"}
	var changes = student2{"John Wick", 2}
	_, err = db.Collection("student").UpdateOne(ctx, selector, bson.M{"$set": changes})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Update success!")
}

func find() {
	db, err := connect1()
	if err != nil {
		log.Fatal(err.Error())
	}

	// bson.M alias dari map[string]interface{}
	csr, err := db.Collection("student").Find(ctx, bson.M{"name": "Wick"})
	if err != nil {
		log.Fatal(err.Error())
	}
	defer csr.Close(ctx)

	result := make([]student2, 0)
	for csr.Next(ctx) {
		var row student2
		err := csr.Decode(&row)
		if err != nil {
			log.Fatal(err.Error())
		}
		result = append(result, row)
	}

	if len(result) > 0 {
		fmt.Println(result)
	}
}

func insert() {
	db, err := connect1()
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = db.Collection("student").InsertOne(ctx, student2{"Wick", 2})
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = db.Collection("student").InsertOne(ctx, student2{"Ethan", 2})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("insert success!")
}

func connect1() (*mongo.Database, error) {
	clientOptions := options.Client()
	clientOptions.ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return client.Database("belajar_golang"), nil
}

var ctx = context.Background()

type student2 struct {
	Name  string `bson:"name"`
	Grade int    `bson:"grade"`
}
