package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/monc"
)

const MomentCollectionName = "moment"

var _ MomentModel = (*customMomentModel)(nil)

type (
	// MomentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMomentModel.
	MomentModel interface {
		momentModel
	}

	customMomentModel struct {
		*defaultMomentModel
	}
)

// NewMomentModel returns a model for the mongo.
func NewMomentModel(url, db string, c cache.CacheConf) MomentModel {
	conn := monc.MustNewModel(url, db, MomentCollectionName, c)
	return &customMomentModel{
		defaultMomentModel: newDefaultMomentModel(conn),
	}
}
