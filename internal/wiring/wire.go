//go:build wireinject
// +build wireinject

//
//go:generate go run github.com/google/wire/cmd/wire
package wiring 

import (
	"github.com/google/wire"

	"goload/internal/configs"
	"goload/internal/dataaccess"
	"goload/internal/handler"
	handlergrpc "goload/internal/handler/grpc"
	"goload/internal/logic"
)

var WireSet = wire.NewSet(
	configs.WireSet,
	dataaccess.WireSet,
	logic.WireSet,
	handler.WireSet,
)

func InitializeGRPCServer(configFilePath configs.ConfigFilePath) (handlergrpc.Server, func(), error) {
	wire.Build(WireSet)
	return nil, nil, nil
}