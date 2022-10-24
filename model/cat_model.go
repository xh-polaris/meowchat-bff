package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/monc"
)

const CatCollectionName = "cat"

var _ CatModel = (*customCatModel)(nil)

type (
	// CatModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCatModel.
	CatModel interface {
		catModel
	}

	customCatModel struct {
		*defaultCatModel
	}
)

// NewCatModel returns a model for the mongo.
func NewCatModel(url, db string, c cache.CacheConf) CatModel {
	conn := monc.MustNewModel(url, db, CatCollectionName, c)
	return &customCatModel{
		defaultCatModel: newDefaultCatModel(conn),
	}
}
