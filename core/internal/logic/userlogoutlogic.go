package logic

import (
	"context"

	"ReactDemoBackend/core/internal/svc"
	"ReactDemoBackend/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogoutLogic {
	return &UserLogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLogoutLogic) UserLogout(req *types.UserLogOutReq) (resp *types.UserLogOutResp, err error) {
	// todo: add your logic here and delete this line

	logx.Info("User Logout:", l.ctx.Value("userId"))
	//如果要做唯一登录，可以在用户表存储一个loginVersion,在这儿把loginVersion +1 ,同时，生成jwtToken的时候也把logingVersion+1,并同时保存到token中
	resp = &types.UserLogOutResp{}
	resp.Code = 200
	return
	return
}
