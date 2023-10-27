//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package di

import (
	"github.com/google/wire"
	"github.com/gsxhnd/tenpai/tenhou/src/service"
)

func InitLog() *service.Log {
	wire.Build(service.NewLog)
	return &service.Log{}
}
