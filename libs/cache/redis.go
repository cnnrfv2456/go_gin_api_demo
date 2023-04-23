package cache

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

type RedisConn struct {
	Password, Host, Port string
	DB                   int
}

var Redis *redis.Client
var ctx = context.Background()

func (t *RedisConn) Connect() {

	addr := fmt.Sprintf("%s:%s", t.Host, t.Port)

	Redis = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: t.Password,
		DB:       t.DB,
	})

	_, err := Redis.Ping(ctx).Result()

	if err != nil {
		log.Panic("Redis init failed: " + err.Error())
	}
}
