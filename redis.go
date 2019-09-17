package main

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis"
)

var (
	redisClient *redis.Client
)

func initRedis() {

	// TODO
	// fix url redis

	redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// try to ping
	_, err := redisClient.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}
}

func getRedis(ctx context.Context, key string) (string, error) {
	return redisClient.Get(key).Result()
}

func setRedis(ctx context.Context, key, value string, ttl int64) error {
	return redisClient.Set(key, value, time.Duration(ttl)*time.Second).Err()
}
