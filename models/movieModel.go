package models

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MovieFormat struct {
	Id   string `bson:"_id"`
	Name string `bson:"name"`
}

func FindAll() []*MovieFormat {
	tempContext := context.TODO()
	cursor, err := moviesCollection.Find(tempContext, bson.M{})

	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(tempContext)

	var movieList []*MovieFormat
	// fmt.Println("Trying to print all movies ")

	for cursor.Next(tempContext) {
		var result bson.M

		if err = cursor.Decode(&result); err != nil {
			log.Fatal(err)
		}
		var movie MovieFormat = MovieFormat{}
		// To convert objectId to hex
		movie.Id = result["_id"].(primitive.ObjectID).Hex()
		movie.Name = result["name"].(string)
		// fmt.Println("Movie")
		// fmt.Println(movie)
		movieList = append(movieList, &movie)

		// fmt.Println(result["_id"].(primitive.ObjectID).Hex())

	}
	return movieList
}
