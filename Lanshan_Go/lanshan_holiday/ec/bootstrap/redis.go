package bootstrap

import (
	"context"
	"demo/ec/global"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

func InitializeRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     global.Ec.Config.Redis.Host + ":" + global.Ec.Config.Redis.Port,
		Password: global.Ec.Config.Redis.Password, // no password set
		DB:       global.Ec.Config.Redis.DB,       // use default DB
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.Ec.Log.Error("Redis connect ping failed, err:", zap.Any("err", err))
		return nil
	}
	return client
}
