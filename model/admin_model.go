package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/monc"
)

const AdminCollectionName = "admin"

var _ AdminModel = (*customAdminModel)(nil)

type (
	// AdminModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAdminModel.
	AdminModel interface {
		adminModel
	}

	customAdminModel struct {
		*defaultAdminModel
	}
)

// NewAdminModel returns a model for the mongo.
func NewAdminModel(url, db string, c cache.CacheConf) AdminModel {
	conn := monc.MustNewModel(url, db, AdminCollectionName, c)
	return &customAdminModel{
		defaultAdminModel: newDefaultAdminModel(conn),
	}
}
