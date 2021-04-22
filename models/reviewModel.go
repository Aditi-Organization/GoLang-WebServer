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
	tempContext := context.TODO()
	movieObjectId, err := primitive.ObjectIDFromHex(movieId)
	fmt.Println(movieObjectId)
	if err != nil {
		log.Fatal(err)
	}

	singleResult := reviewCollection.FindOne(tempContext, bson.M{"_id": movieObjectId})
	if singleResult.Err() != nil {
		log.Fatal(err)
	}
	var result bson.M

	if err = singleResult.Decode(&result); err != nil {
		log.Fatal(err)
	}
	fmt.Println("result ", result)
	return result
	// var movie MovieFormat = MovieFormat{}
	// movie.Id = result["_id"].(primitive.ObjectID).Hex()
	// movie.Name = result["name"].(string)
	// // fmt.Println("Movie")
	// // fmt.Println(movie)
	// movieList = append(movieList, &movie)

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
		// var movie MovieFormat = MovieFormat{}
		// // To convert objectId to hex
		// movie.Id = result["_id"].(primitive.ObjectID).Hex()
		// movie.Name = result["name"].(string)
		// fmt.Println("Movie")
		// fmt.Println(movie)
		// movieList = append(movieList, &movie)
		// fmt.Println(result["_id"].(primitive.ObjectID).Hex())
		fmt.Println(result)
	}
}
