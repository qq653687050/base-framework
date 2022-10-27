package base_data

import (
	"base-framework/inf/conf"
	"base-framework/inf/mysql"
	"base-framework/inf/redisService"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/gorm"
)

/**
 * @description:
 * @author:xy
 * @date:2022/8/30 14:03
 * @Version: 1.0
 */

var ProviderSet = wire.NewSet(NewData, mysql.NewDB, redisService.NewRedis)

type Data struct {
	db *gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, db *gorm.DB) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{db: db}, cleanup, nil
}

type Transaction interface {
	ExecTx(context.Context, func(ctx context.Context) error) error
}

type contextTxKey struct{}

func (d *Data) DB(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(contextTxKey{}).(*gorm.DB)
	if ok {
		return tx
	}
	return d.db
}

func (d *Data) GDB() *gorm.DB {
	return d.db
}

// ExecTx gorm Transaction
func (d *Data) ExecTx(ctx context.Context, fn func(ctx context.Context) error) error {
	return d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, contextTxKey{}, tx)
		return fn(ctx)
	})
}
