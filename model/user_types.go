package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UpdateAt       time.Time          `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt       time.Time          `bson:"createAt,omitempty" json:"createAt,omitempty"`
	Nickname       string             `bson:"nickname" json:"nickname"`
	AvatarUrl      string             `bson:"avatarUrl" json:"avatarUrl"`
	PrivilegeLevel int32              `bson:"privilegeLevel" json:"privilegeLevel"`
}

type UserVO struct {
	ID        primitive.ObjectID `bson:"id,omitempty" json:"id,omitempty"`
	Nickname  string             `bson:"nickname" json:"nickname"`
	AvatarUrl string             `bson:"avatarUrl" json:"avatarUrl"`
}
