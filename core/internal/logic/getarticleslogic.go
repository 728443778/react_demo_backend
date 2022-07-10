package logic

import (
	"context"

	"ReactDemoBackend/core/internal/svc"
	"ReactDemoBackend/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticlesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetArticlesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticlesLogic {
	return &GetArticlesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetArticlesLogic) GetArticles(req *types.GetArticlesReq) (resp *types.GetArticlesResp, err error) {
	// todo: add your logic here and delete this line

	return
}
