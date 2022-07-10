package logic

import (
	"context"

	"ReactDemoBackend/core/helper"
	"ReactDemoBackend/core/internal/svc"
	"ReactDemoBackend/core/internal/types"
	"ReactDemoBackend/core/model"

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
	claims, err := l.svcCtx.JwtParseToken(req.Token)
	resp = &types.GetArticlesResp{}
	if nil != err {
		resp.Code = 401
		resp.Msg = err.Error()
		err = nil
		return
	}
	if helper.GetValueType(claims["userId"]) != "int" {
		resp.Code = 501
		resp.Msg = "storage value invalia type"
	}
	//如果不使用ok且类型不为int时，会导致panic
	v, ok := claims["userId"].(int64)
	if !ok {
		resp.Code = 501
		resp.Msg = "storage value invalia type"
	}

	m := model.NewArticlesModel(l.svcCtx.DbConn)
	artticles, err := m.FindAllByAuthorId(l.ctx, v, 20, req.Page, "id DESC")
	resp.Code = 200
	resp.Msg = ""
	resp.Data = []types.GetArticlesRespData{}
	for _, x := range artticles {
		resp.Data = append(resp.Data, types.GetArticlesRespData{
			Id:        int(x.Id),
			Title:     x.Title,
			Content:   x.Content,
			CreatedAt: x.CreatedAt,
			UpdatedAt: x.UpdatedAt,
		})
	}
	return
}
