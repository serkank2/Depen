package config

import (
	"encoding/json"

	"github.com/go-redis/redis"
)

func Redis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return rdb
}

var RedisClient = Redis()

func SetCache(key string, value interface{}) string {
	data, _ := json.Marshal(value)
	CreateToken := RedisClient.Set(key, data, 0)
	return CreateToken.Val()
}

func GetCache(key string) interface{} {
	val, _ := RedisClient.Get(key).Result()
	// data := config.GetCache(result.Data.Email)
	// var m map[string]interface{}
	// json.Unmarshal([]byte(data.(string)), &m)
	// fmt.Println(m["accesstoken"])
	return val
}
