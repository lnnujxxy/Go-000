package dao

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
)

var (
	ErrNoRows = errors.New("no rows")
)

type Dao struct {
	ctx context.Context
	db  *gorm.DB
	rdb *redis.Client
}

func NewDao(ctx context.Context, db *gorm.DB, rdb *redis.Client) *Dao {
	return &Dao{ctx: ctx, db: db, rdb: rdb}
}
