package rest

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	cache "github.com/victorspringer/http-cache"
	"github.com/victorspringer/http-cache/adapter/memory"
)

var cacheClient *cache.Client

func InitCacheMiddleware(capacity int, ttl time.Duration) error {
	fmt.Println("Init cache middleware: capacity =", capacity, "ttl =", ttl)
	memcached, err := memory.NewAdapter(
		memory.AdapterWithAlgorithm(memory.LRU),
		memory.AdapterWithCapacity(capacity),
	)
	if err != nil {
		return err
	}

	client, err := cache.NewClient(
		cache.ClientWithAdapter(memcached),
		cache.ClientWithTTL(ttl),
		cache.ClientWithRefreshKey("opn"),
	)
	if err != nil {
		return err
	}

	cacheClient = client
	return nil
}

// https://github.com/Columbus-internet/http-cache/blob/9fc5fbc41c27/cache.go#L94
func CacheMiddleware() echo.MiddlewareFunc {
	return echo.WrapMiddleware(cacheClient.Middleware)
}
