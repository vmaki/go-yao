package verifycode

import "go-yao/pkg/redis"

type RedisStore struct {
	RedisClient *redis.Server
	KeyPrefix   string
}

func (s *RedisStore) Set(key string, value string, expiration int64) bool {
	return s.RedisClient.Set(s.KeyPrefix+key, value, expiration)
}

func (s *RedisStore) Get(key string, clear bool) (value string) {
	key = s.KeyPrefix + key
	val := s.RedisClient.Get(key)
	if clear {
		s.RedisClient.Del(key)
	}

	return val
}

func (s *RedisStore) Verify(key, answer string, clear bool) bool {
	v := s.Get(key, clear)

	return v == answer
}
