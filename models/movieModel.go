package models

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MovieFormat struct {
	Id          string `bson:"_id" json:"_id"`
	Name        string `bson:"name" json:"name"`
	Img         string `bson:"img" json:"img"`
	Description string `bson:"description" json:"description"`
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
		var result MovieFormat

		if err = cursor.Decode(&result); err != nil {
			log.Fatal(err)
		}
		// fmt.Println("Movie")
		// fmt.Println(movie)
		movieList = append(movieList, &result)

		// fmt.Println(result["_id"].(primitive.ObjectID).Hex())

	}
	return movieList
}

func FindAllVersion2() []bson.M {
	tempContext := context.TODO()
	cursor, err := moviesCollection.Find(tempContext, bson.M{})

	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(tempContext)

	var movieList []bson.M
	// fmt.Println("Trying to print all movies ")

	for cursor.Next(tempContext) {
		var result bson.M

		if err = cursor.Decode(&result); err != nil {
			log.Fatal(err)
		}
		movieList = append(movieList, result)

		// fmt.Println(result["_id"].(primitive.ObjectID).Hex())

	}
	return movieList
}

func FindMovie(movieName string) MovieFormat {
	tempContext := context.TODO()
	singleResult := moviesCollection.FindOne(tempContext, bson.M{"name": movieName})
	var result MovieFormat

	if err = singleResult.Decode(&result); err != nil {
		log.Fatal(err)
	}
	//fmt.Println("result ", result)
	return result
}

func ShowMovie(Id string) MovieFormat {
	//movieObjectId, err := primitive.ObjectIDFromHex(movieId)

	//	if err != nil {
	//		log.Fatal(err)
	//	}

	tempContext := context.TODO()
	movieId, err := primitive.ObjectIDFromHex(Id)
	singleResult := moviesCollection.FindOne(tempContext, bson.M{"_id": movieId})
	if singleResult.Err() != nil {
		fmt.Println("Error here")
		log.Fatal(err)
	}
	var result MovieFormat

	if err = singleResult.Decode(&result); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Calling result ", result)
	return result
}
