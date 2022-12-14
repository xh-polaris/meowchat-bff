package main

import (
	"flag"
	"fmt"
	"github.com/xh-polaris/meowchat-bff/internal/config"
	"github.com/xh-polaris/meowchat-bff/internal/handler"
	"github.com/xh-polaris/meowchat-bff/internal/svc"
	"github.com/xh-polaris/meowchat-bff/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
	"net/http"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/meowchat.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	httpx.SetErrorHandler(func(err error) (int, interface{}) {
		s, ok := status.FromError(err)
		if ok {
			return http.StatusBadRequest, types.Status{
				Code: int(s.Code()),
				Msg:  s.Message(),
			}
		}
		logx.Error(err.Error())
		return http.StatusInternalServerError, err.Error()
	})

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
