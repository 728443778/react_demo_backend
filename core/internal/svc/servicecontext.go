package svc

import (
	"ReactDemoBackend/core/internal/config"
	"time"

	"github.com/golang-jwt/jwt/v4"
	_ "github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

//如果需要考虑单点登录的问题，可以加入一个版本号控制，用户数据或者redis存储一个用户登录的inc 值，验证的时候，用jwt的版本号值和数据库或者缓存中的值进行判断

type ServiceContext struct {
	Config config.Config
	DbConn sqlx.SqlConn
}

func NewServiceContext(c config.Config) *ServiceContext {
	DbConn := sqlx.NewSqlConn("postgres", c.Mysql.DataSource)
	return &ServiceContext{
		Config: c,
		DbConn: DbConn,
	}
}

func (l *ServiceContext) NewJwtToken(userId int64) (string, error) {
	seconds := l.Config.JwtAuth.AccessExpire
	iat := time.Now().Unix()
	secret := []byte(l.Config.JwtAuth.AccessSecret)
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString(secret)
	// secret := []byte(l.Config.JwtAuth.AccessSecret)
	// nowUnix := time.Now().Unix()
	// claims := JwtClaims{
	// 	userId,
	// 	jwt.StandardClaims{
	// 		ExpiresAt: nowUnix + l.Config.JwtAuth.AccessExpire,
	// 		IssuedAt:  nowUnix,
	// 		Id:        strconv.Itoa(int(userId)),
	// 	},
	// }
	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// return token.SignedString(secret)
}

type JwtClaims struct {
	Uid int64 `json:"uid"`
	jwt.StandardClaims
}

func (l *ServiceContext) JwtParseToken(token string) (jwt.MapClaims, error) {

	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(l.Config.JwtAuth.AccessSecret), nil
	})
	if nil != err {
		return nil, err
	}
	return claims, err
}
