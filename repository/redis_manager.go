package repository

import (
	"github.com/go-redis/redis"
	jsoniter "github.com/json-iterator/go"
)

type RedisManager struct {
	client      redis.Cmdable
}

func NewRedisManager(client redis.Cmdable) *RedisManager {
	return &RedisManager{
		client: client,
	}
}

func (s *RedisManager) remove(key string) error {
	return s.client.Del(key).Err()
}

func (s *RedisManager) insert(key string, value interface{}) error {
	val, _ := jsoniter.Marshal(value)
	return s.client.Set(key, string(val), 0).Err()
}

func (s *RedisManager) insertString(key string, value string) error {
	val, _ := jsoniter.Marshal(value)
	return s.client.Set(key, string(val), 0).Err()
}

func (s *RedisManager) fetch(key string, value interface{}) error {
	val, err := s.client.Get(key).Result()
	if err != nil {
		return err
	}
	return jsoniter.Unmarshal([]byte(val), value)
}

func (s *RedisManager) fetchString(key string) (string, error) {
	return s.client.Get(key).Result()
}
