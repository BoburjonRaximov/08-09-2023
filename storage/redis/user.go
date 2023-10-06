package redis

// import (
// 	"context"
// 	"time"

// 	"github.com/go-redis/cache/v9"
// )

// type cacheRepo struct {
// 	db *cache.Cache
// }

// func NewCache(db *cache.Cache) *cacheRepo {
// 	return &cacheRepo{
// 		db: db,
// 	}
// }

// func (c *cacheRepo) Create(ctx context.Context, key string, obj interface{}, ttl time.Duration) error {
// 	return nil
// }

// func (c *cacheRepo) Get(ctx context.Context, key string, res interface{})(bool, error){
// return true,nil
// }

// func (c *cacheRepo) Delete(ctx context.Context, id string)error{
// 	return nil
// }
