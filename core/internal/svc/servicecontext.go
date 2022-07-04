package svc

import (
	"ReactDemoBackend/core/internal/config"

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
