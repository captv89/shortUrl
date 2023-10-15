package svc

import (
	"shortUrl/api/internal/config"
	"shortUrl/rpc/transform/transformer"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	Transformer transformer.Transformer  // manual code
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
    Transformer: transformer.NewTransformer(zrpc.MustNewClient(c.Transform)), // manual code
	}
}
