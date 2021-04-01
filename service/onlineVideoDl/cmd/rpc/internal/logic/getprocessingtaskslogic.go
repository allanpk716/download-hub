package logic

import (
	"context"

	"github.com/allanpk716/Downloadhub/service/onlineVideoDl/cmd/rpc/internal/svc"
	"github.com/allanpk716/Downloadhub/service/onlineVideoDl/cmd/rpc/onlineVideoDl"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetProcessingTasksLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProcessingTasksLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProcessingTasksLogic {
	return &GetProcessingTasksLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProcessingTasksLogic) GetProcessingTasks(in *onlineVideoDl.TasksInProgressRequest) (*onlineVideoDl.TasksInProgressResponse, error) {
	// todo: add your logic here and delete this line

	return &onlineVideoDl.TasksInProgressResponse{}, nil
}
