package logic

import (
	"context"
	ProxyHub2 "github.com/allanpk716/Downloadhub/service/proxyhub/rpc/ProxyHub"
	svc2 "github.com/allanpk716/Downloadhub/service/proxyhub/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetLogic struct {
	ctx    context.Context
	svcCtx *svc2.ServiceContext
	logx.Logger
}

func NewGetLogic(ctx context.Context, svcCtx *svc2.ServiceContext) *GetLogic {
	return &GetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  获取一个随机代理
func (l *GetLogic) Get(in *ProxyHub2.GetOneReq) (*ProxyHub2.GetOneResp, error) {
	// todo: add your logic here and delete this line

	return &ProxyHub2.GetOneResp{}, nil
}
