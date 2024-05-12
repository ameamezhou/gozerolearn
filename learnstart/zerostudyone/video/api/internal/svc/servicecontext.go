package svc

import (
	"api/internal/config"
	"github.com/zeromicro/go-zero/zrpc"
	"zerostudyone/user/rpc/userclient"
)

type ServiceContext struct {
	Config config.Config
	UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
