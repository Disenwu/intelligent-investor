package service

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func RedisInitialize(cfg *RedisConfig) {
	// 初始化 Redis 连接
	RedisClient = redis.NewClient(&redis.Options{
		Addr:         cfg.Addr,
		Password:     cfg.Password,
		DB:           cfg.Database,
		PoolSize:     cfg.PoolSize,
		MinIdleConns: cfg.MinIdleConns,
	})
	// 测试连接
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := RedisClient.Ping().Result()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect Redis: %v", err))
	}
	fmt.Println("Connected to Redis")
}

type RedisConfig struct {
	Addr         string `mapstructure:"addr"`
	Password     string `mapstructure:"password"`
	Database     int    `mapstructure:"database"`
	PoolSize     int    `mapstructure:"pool-size"`
	MinIdleConns int    `mapstructure:"min-idle-conns"`
}

func NewRedisConfig() *RedisConfig {
	return &RedisConfig{
		Addr:         "127.0.0.1:6379",
		Password:     "",
		Database:     0,
		PoolSize:     100,
		MinIdleConns: 10,
	}
}

func SetAndExpire(key string, value interface{}, expiration time.Duration) error {
	return RedisClient.Set(key, value, expiration).Err()
}

func Get(key string) (string, error) {
	return RedisClient.Get(key).Result()
}

func Delete(key string) error {
	return RedisClient.Del(key).Err()
}
