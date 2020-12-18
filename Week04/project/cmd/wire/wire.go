// +build wireinject
package wire

import (
	"context"
	"github.com/google/wire"
	"gogeekbang/internal/dao"
	"gogeekbang/internal/pkg/config"
	"gogeekbang/internal/pkg/db"
	"gogeekbang/internal/pkg/rdb"
	"gogeekbang/internal/service"
)

func InitService2(ctx context.Context, databaseConfig config.Database, redisConfig config.Redis) service.Service {
	wire.Build(service.NewService, dao.NewDao, db.NewDB, rdb.NewRedis)
	return service.Service{}
}
