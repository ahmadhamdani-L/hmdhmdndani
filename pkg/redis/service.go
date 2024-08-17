package redis

import (
	"lion-super-app/configs"

	redis "github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func Init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     configs.Redis().Host() + ":" + configs.Redis().Port(),
		Password: configs.Redis().Password(), // no password set
		DB:       configs.Redis().Db(),       // use default DB
		PoolSize: configs.Redis().PoolSize(),
	})
}

func IsNil(err error) bool {
	return err == redis.Nil
}
