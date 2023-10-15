package svc

import (
	"shortUrl/rpc/transform/internal/config"
	"shortUrl/rpc/transform/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
  c     config.Config
  Model model.ShorturlModel   // manual code
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		c:             c,
		Model: model.NewShorturlModel(sqlx.NewMysql(c.DataSource), c.Cache), // manual code
	}
}
