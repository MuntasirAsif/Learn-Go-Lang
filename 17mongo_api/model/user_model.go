package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Student struct {
	Id          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	RollNumber  string             `bson:"roll,omitempty" json:"roll,omitempty"`
	PhoneNumber string             `bson:"phoneNumber,omitempty" json:"phoneNumber,omitempty"`
	FullName    string             `bson:"fullName" json:"fullName"`
	FathersName string             `bson:"fathersName" json:"fathersName"`
	MothersName string             `bson:"mothersName" json:"mothersName"`
	Division    string             `bson:"division" json:"division"`
	Result      string             `bson:"result" json:"result"`
	Gender      string             `bson:"gender" json:"gender"`
}
