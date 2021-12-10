// Package initialize
// @program: arco-design-pro-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2021-12-10 15:09
package initialize

import (
	"context"
	"fmt"

	"gin-basice/global"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

func Redis() {
	redisCfg := global.AdpConfig.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping(context.TODO()).Result()
	if err != nil {
		global.AdpLog.Error("redis connect ping failed, err:", zap.Error(err))
		panic(fmt.Sprintf("redis connect ping failed, err: %s", err.Error()))
	} else {
		global.AdpLog.Info("redis connect ping response:", zap.String("pong", pong))
		global.AdpRedis = client
	}
}
