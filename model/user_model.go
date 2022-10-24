package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/monc"
)

const UserCollectionName = "user"

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
	}

	customUserModel struct {
		*defaultUserModel
	}
)

// NewUserModel returns a model for the mongo.
func NewUserModel(url, db string, c cache.CacheConf) UserModel {
	conn := monc.MustNewModel(url, db, UserCollectionName, c)
	return &customUserModel{
		defaultUserModel: newDefaultUserModel(conn),
	}
}
