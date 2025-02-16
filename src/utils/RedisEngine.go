package utils

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
)

func RedisConnection() *redis.Client {
	conn, err := getConnection()

	if err != nil {
		log.Default().Println("redis 连接失败")
		return nil
	}

	return conn
}

func getConnection() (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Ping(context.Background()).Err()

	// 连接失败
	if err != nil {
		log.Default().Println("连接失败", err)
		return nil, err
	}

	return rdb, nil
}

var DefaultRedis = RedisConnection()
