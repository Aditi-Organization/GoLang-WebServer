package models

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserDetailFormat struct {
	userId   string `bson:"u_id"`
	userName string `bson:"u_name"`
	userEmail string `bson:"u_email"`
	userPassword string `bson:"u_password"`

}

func FindAllUsers() []*UserDetailFormat {
	tempContext := context.TODO()
	cursor, err := userCollection.Find(tempContext, bson.M{})

	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(tempContext)

	var userList []*UserDetailFormat
	// fmt.Println("Trying to print all users ")
	for cursor.Next(tempContext) {
		var result bson.M

		if err = cursor.Decode(&result); err != nil {
			log.Fatal(err)
		}
		var user UserDetailFormat = UserDetailFormat{}
		// To convert objectId to hex
		user.userId = result["u_id"].(primitive.ObjectID).Hex()
		user.userName = result["u_name"].(string)
		user.userEmail = result["u_email"].(string)
		user.userPassword = result["u_password"].(string)
		
		userList = append(userList, &user)

	}
	return userList
}
