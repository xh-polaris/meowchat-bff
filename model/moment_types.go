package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Comment struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	User     UserVO             `bson:"user" json:"user"`
	UpdateAt time.Time          `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt time.Time          `bson:"createAt,omitempty" json:"createAt,omitempty"`
}

type Moment struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UpdateAt  time.Time          `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt  time.Time          `bson:"createAt,omitempty" json:"createAt,omitempty"`
	CatId     string             `bson:"catId" json:"catId"`
	Photos    []string           `bson:"photos" json:"photos"`
	Status    int32              `bson:"status" json:"status"`
	Text      string             `bson:"text" json:"text"`
	User      UserVO             `bson:"user" json:"user"`
	UserLikes int32              `bson:"userLikes" json:"userLikes"`
}
