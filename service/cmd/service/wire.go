//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"service/internal/biz"
	"service/internal/conf"
	"service/internal/data"
	"service/internal/server"
	"service/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init 这里wireApp中参数的作用是把config.yaml文件中的数据内容导入internal/config/config.pb.go中
func wireApp(*conf.Server, *conf.Data, *conf.Wechat, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
