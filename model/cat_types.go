package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Cat struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UpdateAt     time.Time          `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt     time.Time          `bson:"createAt,omitempty" json:"createAt,omitempty"`
	Age          string             `bson:"age" json:"age"`
	CampusId     int32              `bson:"campusId" json:"campusId"`
	ColorId      int32              `bson:"colorId" json:"colorId"`
	Details      string             `bson:"details" json:"details"`
	Name         string             `bson:"name" json:"name"`
	Popularity   int32              `bson:"popularity" json:"popularity"`
	Sex          string             `bson:"sex" json:"sex"`
	StatusId     int32              `bson:"statusId" json:"statusId"`
	Area         string             `bson:"area" json:"area"`
	IsSnipped    bool               `bson:"isSnipped" json:"isSnipped"`
	IsSterilized bool               `bson:"isSterilized" json:"isSterilized"`
	AvatarUrl    string             `bson:"avatarUrl" json:"avatarUrl"`
}
