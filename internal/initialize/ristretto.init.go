package initialize

import (
	"github.com/anle/codebase/global"
	"github.com/dgraph-io/ristretto/v2"
)

func InitRistretto() {
	localCache, err := ristretto.NewCache(&ristretto.Config[string, string]{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})
	if err != nil {
		panic(err)
	}
	global.LocalCache = localCache
}
