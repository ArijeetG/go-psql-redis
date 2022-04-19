package helper

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

func GetCacheClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "password",
		DB:       0,
	})
	return rdb
}

func SetCache(rdb *redis.Client, ctx context.Context, key string, value []byte) int {
	_, err := rdb.Set(ctx, key, value, 5*time.Minute).Result()
	if err != nil {
		log.Println(err.Error())
		return 0
	}
	return 1
}

func GetCache(rdb *redis.Client, ctx context.Context, key string) interface{} {
	result, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return true // data not present in cache
	} else if err != nil {
		log.Println(err.Error())
		return false // something went wrong
	} else {
		return result //present in cache
	}
}
