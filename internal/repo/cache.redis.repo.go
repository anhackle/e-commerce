package repo

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/bsm/redislock"
	"github.com/redis/go-redis/v9"
)

type IRedisCache interface {
	Get(ctx context.Context, key string) (value string, err error)
	Set(ctx context.Context, key string, value interface{}, expiration int64) (err error)
	Del(ctx context.Context, key string) error
	Incr(ctx context.Context, key string) (int64, error)
	Decr(ctx context.Context, key string) (int64, error)
	Exists(ctx context.Context, key string) (bool, error)

	WithDistributedLock(ctx context.Context, key string, ttlSeconds int, fn func(ctx context.Context) error) error
}

type redisCache struct {
	client *redis.Client
	locker *redislock.Client
}

// WithDistributedLock implements IRedisCache.
func (r *redisCache) WithDistributedLock(ctx context.Context, key string, ttlSeconds int, fn func(ctx context.Context) error) error {
	lockTTL := time.Duration(ttlSeconds) * time.Second
	lock, err := r.locker.Obtain(ctx, key, lockTTL, nil)
	if err == redislock.ErrNotObtained {
		return errors.New("lock can not obtained")
	} else if err != nil {
		return errors.New("fail to obtain lock")
	}

	defer lock.Release(ctx)

	return fn(ctx)
}

// Decr implements IRedisCache.
func (r *redisCache) Decr(ctx context.Context, key string) (int64, error) {
	panic("unimplemented")
}

// Del implements IRedisCache.
func (r *redisCache) Del(ctx context.Context, key string) error {
	panic("unimplemented")
}

// Exists implements IRedisCache.
func (r *redisCache) Exists(ctx context.Context, key string) (bool, error) {
	panic("unimplemented")
}

// Get implements IRedisCache.
func (r *redisCache) Get(ctx context.Context, key string) (value string, err error) {
	value, err = r.client.Get(ctx, key).Result()
	if err != nil && err != redis.Nil {
		return value, err
	}

	return value, nil
}

// Incr implements IRedisCache.
func (r *redisCache) Incr(ctx context.Context, key string) (int64, error) {
	panic("unimplemented")
}

// Set implements IRedisCache.
func (r *redisCache) Set(ctx context.Context, key string, value any, expiration int64) (err error) {
	jsonValue, err := json.Marshal(value)
	if err != nil {
		return err
	}

	err = r.client.Set(ctx, key, jsonValue, time.Duration(expiration)*time.Second).Err()
	if err != nil {
		return err
	}

	return nil
}

func NewRedisCache(client *redis.Client) IRedisCache {
	return &redisCache{
		client: client,
		locker: redislock.New(client),
	}
}
