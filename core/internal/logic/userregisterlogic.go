package logic

import (
	"context"
	"database/sql"
	"strings"

	"ReactDemoBackend/core/helper"
	"ReactDemoBackend/core/internal/svc"
	"ReactDemoBackend/core/internal/types"
	"ReactDemoBackend/core/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterReq) (resp *types.UserRegisterResp, err error) {
	// todo: add your logic here and delete this line
	userName := req.UserName
	password := req.Password
	resp = &types.UserRegisterResp{}
	if len(strings.TrimSpace(userName)) == 0 || len(strings.TrimSpace(password)) == 0 {
		resp.Code = 406
		resp.Msg = "参数错误"
		return
	}
	//是否已经注册
	userModel := model.NewUsersModel(l.svcCtx.DbConn)
	_, err = userModel.FindOneByUserName(l.ctx, userName)
	if nil == err {
		resp.Code = 406
		resp.Msg = "用户已存在"
		return
	}
	user := model.Users{}
	user.Status = sql.NullInt64{}
	user.Status.Int64 = 1
	user.Username = userName
	tmpPass, err := helper.HashPassword(password)
	if nil != err {
		resp.Code = 505
		resp.Msg = "generate password failed"
		return
	}
	user.Password = string(tmpPass)
	result, err := userModel.Insert(l.ctx, &user)
	if nil != err {
		resp.Code = 409
		resp.Msg = err.Error()
		return
	}
	rowEffected, _ := result.RowsAffected()
	if rowEffected <= 0 {
		resp.Code = 501
		resp.Msg = "Register failed"
		return
	}
	resp.Code = 200
	resp.Msg = "success"
	return
}
