package redisService

import (
	"base-framework/inf/conf"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

/**
 * @description:
 * @author:xy
 * @date:2022/8/30 16:49
 * @Version: 1.0
 */
var ctx = context.Background()

func NewRedis(c *conf.Data) redis.UniversalClient {
	if c.Redis.Enable {
		//redis 通用客户端
		client := redis.NewUniversalClient(&redis.UniversalOptions{
			Addrs:              c.Redis.Addr,
			Password:           c.Redis.Password,
			DB:                 int(c.Redis.Db),
			PoolSize:           int(c.Redis.Pool.Size),
			MinIdleConns:       int(c.Redis.Pool.MinIdleConns),
			DialTimeout:        c.Redis.DialTimeout.AsDuration(),
			ReadTimeout:        c.Redis.ReadTimeout.AsDuration(),
			WriteTimeout:       c.Redis.WriteTimeout.AsDuration(),
			PoolTimeout:        c.Redis.Pool.Timeout.AsDuration(),
			IdleCheckFrequency: c.Redis.Pool.IdleCheckFrequency.AsDuration(),
			IdleTimeout:        c.Redis.Pool.IdleTimeout.AsDuration(),
			MaxConnAge:         c.Redis.Pool.MaxConnAge.AsDuration(),

			//钩子函数
			OnConnect: func(ctx context.Context, conn *redis.Conn) error {
				//仅当客户端执行命令时需要从连接池获取连接时，如果连接池需要新建连接时则会调用此钩子函数
				fmt.Printf("conn=%v\n", conn)
				return nil
			},
		})

		return client
	}
	return nil
}
