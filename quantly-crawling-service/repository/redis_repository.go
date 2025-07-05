package repository

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

func PostDataToRedisDB[T any](rdb *redis.Client, data *T, key string, ctx context.Context) error {
	_, err := rdb.JSONSet(ctx, key, "$", data).Result()
	rdb.Expire(ctx, key, 2*time.Hour)
	if err != nil {
		return err
	}
	return nil
}
