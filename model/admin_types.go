package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Admin struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UpdateAt time.Time          `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt time.Time          `bson:"createAt,omitempty" json:"createAt,omitempty"`
	Name     string             `bson:"name" json:"name"`
	Phone    string             `bson:"phone" json:"phone"`
	Wechat   string             `bson:"wechat" json:"wechat"`
	Position string             `bson:"position" json:"position"`
	Status   int64              `bson:"status" json:"status"`
}
