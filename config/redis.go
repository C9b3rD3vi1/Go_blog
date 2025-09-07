package config

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)


func TestRedisConnection() error {
    client := redis.NewClient(&redis.Options{
        Addr:     "127.0.0.1:6379",
        Password: "",
        DB:       0,
    })

    ctx := context.Background()

    // Test connection
    pong, err := client.Ping(ctx).Result()
    if err != nil {
        log.Printf("❌ Redis connection failed: %v", err)
        return err
    }
    log.Printf("✅ Redis Ping: %s", pong)

    // Test write/read
    testKey := "test_session_key"
    testValue := "test_value_" + time.Now().Format("150405")

    err = client.Set(ctx, testKey, testValue, 10*time.Second).Err()
    if err != nil {
        log.Printf("❌ Redis set failed: %v", err)
        return err
    }

    retrieved, err := client.Get(ctx, testKey).Result()
    if err != nil {
        log.Printf("❌ Redis get failed: %v", err)
        return err
    }

    log.Printf("✅ Redis test successful: set '%s'='%s', got '%s'", testKey, testValue, retrieved)

    return client.Close()
}
