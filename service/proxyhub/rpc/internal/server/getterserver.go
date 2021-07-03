// Code generated by goctl. DO NOT EDIT!
// Source: proxyhub.proto

package server

import (
	"context"
	ProxyHub2 "github.com/allanpk716/Downloadhub/service/proxyhub/rpc/ProxyHub"
	logic2 "github.com/allanpk716/Downloadhub/service/proxyhub/rpc/internal/logic"
	svc2 "github.com/allanpk716/Downloadhub/service/proxyhub/rpc/internal/svc"
)

type GetterServer struct {
	svcCtx *svc2.ServiceContext
}

func NewGetterServer(svcCtx *svc2.ServiceContext) *GetterServer {
	return &GetterServer{
		svcCtx: svcCtx,
	}
}

//  获取一个随机代理
func (s *GetterServer) Get(ctx context.Context, in *ProxyHub2.GetOneReq) (*ProxyHub2.GetOneResp, error) {
	l := logic2.NewGetLogic(ctx, s.svcCtx)
	return l.Get(in)
}