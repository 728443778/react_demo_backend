package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Mysql struct {
		DataSource string
	}

	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}
}
