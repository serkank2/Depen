package config

import (
	"fmt"

	"github.com/go-redis/redis"
)

func Redis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	fmt.Println("Redis Connected")
	return rdb
}

var RedisClient = Redis()

func SetCache(key string, value interface{}) {
	RedisClient.Set(key, value, 0)
}

func GetCache(key string) interface{} {
	val, _ := RedisClient.Get(key).Result()
	return val
}
