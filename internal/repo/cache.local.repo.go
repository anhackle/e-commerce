package repo

import (
	"context"
	"encoding/json"
	"time"

	"github.com/dgraph-io/ristretto/v2"
)

type ILocalCache interface {
	Get(ctx context.Context, key string) (value string, found bool)
	Set(ctx context.Context, key string, value interface{}) (result bool)
	SetWithTTL(ctx context.Context, key string, value interface{}) (result bool, err error)
	Del(ctx context.Context, key string) error
	Incr(ctx context.Context, key string) (int64, error)
	Decr(ctx context.Context, key string) (int64, error)
	Exists(ctx context.Context, key string) (bool, error)
}

type localCache struct {
	client *ristretto.Cache[string, string]
}

// SetWithTTL implements ILocalCache.
func (lc *localCache) SetWithTTL(ctx context.Context, key string, value interface{}) (result bool, err error) {
	jsonValue, err := json.Marshal(value)
	if err != nil {
		return result, err
	}

	return lc.client.SetWithTTL(key, string(jsonValue), 1, 5*time.Minute), nil
}

// Decr implements IRedisCache.
func (lc *localCache) Decr(ctx context.Context, key string) (int64, error) {
	panic("unimplemented")
}

// Del implements IRedisCache.
func (lc *localCache) Del(ctx context.Context, key string) error {
	panic("unimplemented")
}

// Exists implements IRedisCache.
func (lc *localCache) Exists(ctx context.Context, key string) (bool, error) {
	panic("unimplemented")
}

// Get implements IRedisCache.
func (lc *localCache) Get(ctx context.Context, key string) (value string, found bool) {
	value, found = lc.client.Get(key)

	return value, found
}

// Incr implements IRedisCache.
func (lc *localCache) Incr(ctx context.Context, key string) (int64, error) {
	panic("unimplemented")
}

// Set implements IRedisCache.
func (lc *localCache) Set(ctx context.Context, key string, value any) (result bool) {
	jsonValue, _ := json.Marshal(value)

	return lc.client.Set(key, string(jsonValue), 1)
}

func NewLocalCache(client *ristretto.Cache[string, string]) ILocalCache {
	return &localCache{
		client: client,
	}
}
