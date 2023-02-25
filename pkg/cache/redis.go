package cache

import (
	"go-yao/common/global"
	"go-yao/pkg/redis"
)

// RedisStore 实现 cache.Store interface
type RedisStore struct {
	RedisClient *redis.Server
	KeyPrefix   string
}

func NewRedisStore(address string, username string, password string, db int) *RedisStore {
	rs := &RedisStore{}
	rs.RedisClient = redis.Connect(address, username, password, db)
	rs.KeyPrefix = global.Conf.Application.Name + ":cache:"

	return rs
}

func (s *RedisStore) Set(key string, value string, expireTime int64) bool {
	return s.RedisClient.Set(s.KeyPrefix+key, value, expireTime)
}

func (s *RedisStore) Get(key string) string {
	return s.RedisClient.Get(s.KeyPrefix + key)
}

func (s *RedisStore) Has(key string) bool {
	return s.RedisClient.Has(s.KeyPrefix + key)
}

func (s *RedisStore) Del(key string) {
	s.RedisClient.Del(s.KeyPrefix + key)
}
