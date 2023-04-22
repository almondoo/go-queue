package redis

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

var conn *redis.Client

func init() {
	conn = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

// redisから取得する
func Get(ctx context.Context, key string) ([]byte, error) {
	return conn.Get(ctx, key).Bytes()
}

// redisに保存する
func Set(ctx context.Context, value interface{}) error {
	uuid, _ := uuid.NewRandom()
	value, _ = json.Marshal(value)
	return conn.Set(ctx, uuid.String(), value, 0).Err()
}

func AllGet(ctx context.Context) {
	// foreachで回して全取得したデータを返す
	data := conn.Keys(ctx, "*").Val()
	fmt.Println(data)
}
