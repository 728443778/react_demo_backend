// Code generated by goctl. DO NOT EDIT.
package types

type UserRegisterReq struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type UserLogOutReq struct {
	Token string `json:"token"`
}

type UserRegisterResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type UserLogOutResp struct {
	Code int `json:"code"`
}

type LoginRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Code int               `json:"code"`
	Data LoginResponseData `json:"data"`
	Msg  string            `json:"msg"`
}

type LoginResponseData struct {
	Avatar    string `json:"avatar"`
	UserName  string `json:"username"`
	Token     string `json:"token"`
	ExpiredAt int64    `json:"exp"`
}
