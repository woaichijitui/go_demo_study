package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gozero_mall_study/service/product/model"
	"gozero_mall_study/service/product/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config

	ProductModel model.ProductModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:       c,
		ProductModel: model.NewProductModel(conn, c.CacheRedis),
	}
}
