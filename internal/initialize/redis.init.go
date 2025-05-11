package initialize

import (
	"context"
	"fmt"

	"github.com/anle/codebase/global"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var ctx = context.Background()

func InitRedis() {
	r := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", r.Host, r.Port),
		Password: r.Password,
		DB:       r.Database,
		PoolSize: 10,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Error("Redis initialization error", zap.Error(err))
	}

	global.Rdb = rdb
}

func InitRedisSentinel() {
	rdb := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName: "redis-master",
		SentinelAddrs: []string{
			"sentinel-master:26379",
			"sentinel-slave1:26379",
			"sentinel-slave2:26379",
		},
		DB:       0,
		Password: "",
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Error("Redis initialization error", zap.Error(err))
	}

	err = rdb.Set(ctx, "test_key", "Hello Redis Sentinel", 0).Err()
	if err != nil {
		global.Logger.Error("Redis set error", zap.Error(err))
		panic(err)
	}

	global.Rdb = rdb
}
