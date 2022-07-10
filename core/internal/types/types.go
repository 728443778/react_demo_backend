// Code generated by goctl. DO NOT EDIT.
package types

type GetArticlesReq struct {
	Token string `form:"token"`
	Page  int    `form:"page"`
}

type GetArticlesResp struct {
	Code int                   `json:"code"`
	Msg  string                `json:"msg"`
	Data []GetArticlesRespData `json:"data"`
}

type GetArticlesRespData struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

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
	ExpiredAt int64  `json:"exp"`
}
