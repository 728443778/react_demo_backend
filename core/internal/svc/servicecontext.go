package svc

import (
	"ReactDemoBackend/core/internal/config"

	"github.com/golang-jwt/jwt/v4"
	_ "github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

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

func (l *ServiceContext) NewJwtToken(secret string, iat int64, seconds int64, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secret))
}
