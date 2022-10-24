package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Carousel struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UpdateAt time.Time          `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt time.Time          `bson:"createAt,omitempty" json:"createAt,omitempty"`
	ImgUrl   string             `bson:"imgUrl" json:"imgUrl"`
	Order    int32              `bson:"order" json:"order"`
	SrcUrl   string             `bson:"srcUrl" json:"srcUrl"`
	TypeId   int32              `bson:"typeId" json:"typeId"`
}
