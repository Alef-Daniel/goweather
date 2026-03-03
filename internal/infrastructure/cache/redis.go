package cache

import (
	"context"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	client *redis.Client
}

func (r *Redis) Set(ctx context.Context, key, value string, exp time.Duration) error {
	return r.client.Set(ctx, key, value, exp).Err()

}

func (r *Redis) Get(ctx context.Context, key string) (string, error) {
	result, err := r.client.Get(ctx, key).Bytes()
	if errors.Is(err, redis.Nil) {
		return "", nil
	}
	if err != nil {
		return "", err
	}
	return string(result), nil
}
func (r Redis) Delete(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}

func NewRedis(addr, pass string) *Redis {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
	})

	return &Redis{client: rdb}
}
