package repository

import (
	"fmt"
	"ir.safari.shortlink/model"
)

type CachedUrlRedisRepository interface {
	Insert(int64, *model.CachedUrl) error
}

type cachedUrlRedisRepository struct {
	redisManager *RedisManager
}

const prefix = "cached_url_keys_%d"

func (c cachedUrlRedisRepository) Insert(key int64, value *model.CachedUrl) error {
	fkey := fmt.Sprintf(prefix, key)
	return c.redisManager.insert(fkey, value)
}

func NewCachedUrlRedisRepository(redisManager *RedisManager) CachedUrlRedisRepository {
	return &cachedUrlRedisRepository{
		redisManager: redisManager,
	}
}