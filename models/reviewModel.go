package models

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// type Review struct {
// 	id          string
// 	movieId     string
// 	description string
// 	rating      int
// 	userId      string
// }

func GetReview(movieId string) bson.M {
	movieObjectId, err := primitive.ObjectIDFromHex(movieId)

	if err != nil {
		log.Fatal(err)
	}

	tempContext := context.TODO()
	singleResult := reviewCollection.FindOne(tempContext, bson.M{"_id": movieObjectId})
	if singleResult.Err() != nil {
		log.Fatal(err)
	}
	var result bson.M

	if err = singleResult.Decode(&result); err != nil {
		log.Fatal(err)
	}
	// fmt.Println("result ", result)
	return result
}

func GetMovieReviews(movieId string) []bson.M {
	movieObjectId, err := primitive.ObjectIDFromHex(movieId)

	if err != nil {
		log.Fatal(err)
	}

	tempContext := context.TODO()
	cursor, err := reviewCollection.Find(tempContext, bson.M{"movieId": movieObjectId})

	var movieReviews []bson.M
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(tempContext)

	for cursor.Next(tempContext) {
		var result bson.M

		if err = cursor.Decode(&result); err != nil {
			log.Fatal(err)
		}
		// fmt.Println(result)
		movieReviews = append(movieReviews, result)
	}
	return movieReviews
}

func PrintAllReviews() {
	tempContext := context.TODO()
	cursor, err := reviewCollection.Find(tempContext, bson.M{})

	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(tempContext)

	for cursor.Next(tempContext) {
		var result bson.M

		if err = cursor.Decode(&result); err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
	}
}
