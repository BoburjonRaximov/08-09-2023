package redis

import (
	"context"
	"fmt"
	"playground/cpp-bootcamp/config"
	"playground/cpp-bootcamp/storage"

	"github.com/go-redis/cache/v9"
	goRedis "github.com/redis/go-redis/v9"
)

type cacheStrg struct {
	db     *cache.Cache
	cacheR *cacheRepo
	
}

func NewCache(ctx context.Context, cfg config.Config) (storage.CacheI, error) {
	redisClient := goRedis.NewClient(&goRedis.Options{
		Addr:     fmt.Sprintf("%v:%v", cfg.RedisHost, cfg.RedisPort),
		Password: cfg.RedisPassword,
		DB:       0,
	})
	redisCache := cache.New(&cache.Options{
		Redis: redisClient,
	})
	return &cacheStrg{
		db: redisCache,
	}, nil
}

func (d *cacheStrg) Cache() storage.RedisI {
	if d.cacheR == nil {
		d.cacheR = NewCacheRepo(d.db)
	}
	return d.cacheR
}
