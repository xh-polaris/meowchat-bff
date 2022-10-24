package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/monc"
)

const CarouselCollectionName = "carousel"

var _ CarouselModel = (*customCarouselModel)(nil)

type (
	// CarouselModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCarouselModel.
	CarouselModel interface {
		carouselModel
	}

	customCarouselModel struct {
		*defaultCarouselModel
	}
)

// NewCarouselModel returns a model for the mongo.
func NewCarouselModel(url, db string, c cache.CacheConf) CarouselModel {
	conn := monc.MustNewModel(url, db, CarouselCollectionName, c)
	return &customCarouselModel{
		defaultCarouselModel: newDefaultCarouselModel(conn),
	}
}
