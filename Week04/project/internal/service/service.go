package service

import (
	"context"
	"gogeekbang/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func NewService(ctx context.Context, dao *dao.Dao) Service {
	return Service{
		ctx: ctx,
		dao: dao,
	}
}
