// Code generated by goctl. DO NOT EDIT.
// Source: rpctest.proto

package rpctestclient

import (
	"context"

	"rpctest/rpctest"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Request  = rpctest.Request
	Response = rpctest.Response

	Rpctest interface {
		Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	}

	defaultRpctest struct {
		cli zrpc.Client
	}
)

func NewRpctest(cli zrpc.Client) Rpctest {
	return &defaultRpctest{
		cli: cli,
	}
}

func (m *defaultRpctest) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := rpctest.NewRpctestClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}
