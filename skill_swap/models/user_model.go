package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	First_name    *string             `bson:"first_name" json:"first_name" validate: required, min=2,max=100`
	Last_name     *string             `bson:"last_name" json:"last_name"`
	Email         *string             `bson:"email" json:"email"`
	Phone         *string             `bson:"phone" json:"phone"`
	Token         *string             `bson:"token" json:"token"`
	User_type     *string             `bson:"user_type" json:"user_type"`
	Refresh_token *string             `bson:"refresh_token" json:"refresh_token"`
	Created_at    *string             `bson:"created_at" json:"created_at"`
	Updated_at    *string             `bson:"updated_at" json:"updated_at"`
	User_id       *string             `bson:"user_id" json:"user_id"`
}
