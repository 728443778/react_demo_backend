package logic

import (
	"context"
	"time"

	"ReactDemoBackend/core/helper"
	"ReactDemoBackend/core/internal/svc"
	"ReactDemoBackend/core/internal/types"
	"ReactDemoBackend/core/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLogic) User(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	// todo: add your logic here and delete this line
	userModel := model.NewUsersModel(l.svcCtx.DbConn)
	findUser, err := userModel.FindOneByUserName(l.ctx, req.UserName)
	resp = &types.LoginResponse{}
	if nil != err {
		logx.Error(err)
		resp.Code = 406
		resp.Msg = "用户不存在或者密码错误"
		return
	}

	err = helper.CheckPassword(req.Password, findUser.Password)
	if nil != err {
		logx.Error(err)

		resp.Code = 406
		resp.Msg = "用户不存在或者密码错误"
		return
	}
	toekn, err := l.svcCtx.NewJwtToken(findUser.Id)
	if nil != err {
		logx.Error(err)
		resp.Code = 500
		resp.Msg = "sever error"
	}
	resp.Code = 200
	resp.Data = types.LoginResponseData{
		Avatar:    "",
		Token:     toekn,
		UserName:  findUser.Username,
		ExpiredAt: time.Now().Unix() + l.svcCtx.Config.JwtAuth.AccessExpire,
	}
	return
}
