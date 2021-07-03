package logic

import (
	"context"
	svc2 "github.com/allanpk716/Downloadhub/service/proxyhub/api/internal/svc"
	types2 "github.com/allanpk716/Downloadhub/service/proxyhub/api/internal/types"
	getter2 "github.com/allanpk716/Downloadhub/service/proxyhub/rpc/getter"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc2.ServiceContext
}

func NewGetLogic(ctx context.Context, svcCtx *svc2.ServiceContext) GetLogic {
	return GetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLogic) Get(req types2.GetOneReq) (*types2.GetOneResp, error) {

	resp, err := l.svcCtx.Getter.Get(l.ctx, &getter2.GetOneReq{
		UseUrl: req.UseUrl,
		Stable: req.Stable,
	})
	if err != nil {
		return nil, err
	}

	return &types2.GetOneResp{
		ID: resp.ID,
		IP: resp.IP,
	}, nil
}
