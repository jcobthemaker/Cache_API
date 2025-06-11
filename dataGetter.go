package main

import (
	"context"
	"fmt"
	"os"
	"time"
	"github.com/go-redis/redis/v8"
	"cache_API/db"
	"database/sql"
	_ "github.com/lib/pq"
)

var redisClient *redis.Client

type RedisConfig struct {
    Addr     string
    DB       int
}

func InitRedis() *redis.Client {
    cfg := loadRedisConfig()

    redisClient = redis.NewClient(&redis.Options{
        Addr:     cfg.Addr,
        DB:       cfg.DB,
    })

    return redisClient
}

func loadRedisConfig() *RedisConfig {
    return &RedisConfig{
        Addr:     os.Getenv("REDIS_ADDR"),
        DB:       0,                            
    }
}


func GetOrSaveRecord(ctx context.Context, database *sql.DB, key string, value string) (string, error) {
    val, err := redisClient.Get(ctx, key).Result()
    if err == nil {
        fmt.Println("Found in Redis:", key, val)
    }
    
    if err != redis.Nil {
        fmt.Printf("Redis error: %v\n", err)
        return "", err 
    }

    val, err = db.Get(ctx, database, key)
    if err != nil {
        fmt.Printf("DB Get error: %v\n", err)
        if err != nil && err != sql.ErrNoRows {
            return "", err
        }
    }
    if val != "" {
        fmt.Println("Found in DB:", key, val)
        if err := redisClient.Set(ctx, key, val, 5*time.Minute).Err(); err != nil {
            fmt.Printf("Warning: failed to cache in Redis: %v\n", err)
        }
        return val, nil
    }

    fmt.Println("Inserting into DB:", key, value)
    if err := db.Set(ctx, database, key, value); err != nil {
        fmt.Printf("DB Set error: %v\n", err)
        return "", err
    }

    if err := redisClient.Set(ctx, key, value, 5*time.Minute).Err(); err != nil {
        fmt.Printf("Warning: failed to cache in Redis: %v\n", err)
    }

    return value, nil
}


func GetAllCached(ctx context.Context) map[string]string {
	valueMap := make(map[string]string)

	cacheIter := redisClient.Scan(ctx, 0, "*", 0).Iterator()

	for cacheIter.Next(ctx) {
		key := cacheIter.Val()
		val, err := redisClient.Get(ctx, key).Result()
		if err == redis.Nil {
			continue
		} else if err != nil {
			continue
		}

		valueMap[key] = val
	}

	return valueMap
}
