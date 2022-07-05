package main

import (
	"flag"
	"fmt"
	"net/http"

	"ReactDemoBackend/core/internal/config"
	"ReactDemoBackend/core/internal/handler"
	"ReactDemoBackend/core/internal/svc"

	"ReactDemoBackend/core/middleware"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithCustomCors(nil, nil, ""))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)

	// 全局中间件
	corsMiddlewre := middleware.NewCorsMiddleware()
	server.Use(corsMiddlewre.Handle)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

func Cors(w http.ResponseWriter) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET, PUT, DELETE, POST")
	w.Header().Add("Access-Control-Allow-Credentail", "true")

}
