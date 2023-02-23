package redis

import (
	"context"
	"go-yao/pkg/logger"
	"sync"
	"time"

	redisLib "github.com/redis/go-redis/v9"
)

type Server struct {
	Context context.Context
	Client  *redisLib.Client
}

var (
	once   sync.Once
	Client *Server
)

func ConnectRedis(address string, username string, password string, db int) {
	once.Do(func() {
		Client = Connect(address, username, password, db)
	})
}

func Connect(address string, username string, password string, db int) *Server {
	rds := &Server{}
	rds.Context = context.Background()
	rds.Client = redisLib.NewClient(&redisLib.Options{
		Addr:     address,
		Username: username,
		Password: password,
		DB:       db,
	})

	err := rds.Ping()
	if err != nil {
		panic("Redis connection failure, err: " + err.Error())
	}

	return rds
}

func (rds Server) Ping() error {
	_, err := rds.Client.Ping(rds.Context).Result()
	return err
}

func (rds Server) Set(key string, value interface{}, expiration int64) bool {
	if err := rds.Client.Set(rds.Context, key, value, time.Duration(expiration)*time.Second).Err(); err != nil {
		logger.ErrorString("Redis", "Set", err.Error())
		return false
	}

	return true
}

func (rds Server) Get(key string) string {
	if result, err := rds.Client.Get(rds.Context, key).Result(); err != nil {
		if err != redisLib.Nil {
			logger.ErrorString("Redis", "Get", err.Error())
		}

		return ""
	} else {
		return result
	}
}

// Has 判断一个 key 是否存在，内部错误和 redis.Nil 都返回 false
func (rds Server) Has(key string) bool {
	if _, err := rds.Client.Get(rds.Context, key).Result(); err != nil {
		if err != redisLib.Nil {
			logger.ErrorString("Redis", "Has", err.Error())
		}

		return false
	}

	return true
}

func (rds Server) Del(keys ...string) bool {
	if err := rds.Client.Del(rds.Context, keys...).Err(); err != nil {
		logger.ErrorString("Redis", "Del", err.Error())
		return false
	}

	return true
}

func (rds Server) Incr(key string) bool {
	if err := rds.Client.Incr(rds.Context, key).Err(); err != nil {
		logger.ErrorString("Redis", "Incr", err.Error())
		return false
	}

	return true
}

func (rds Server) IncrBy(key string, value int64) bool {
	if err := rds.Client.IncrBy(rds.Context, key, value).Err(); err != nil {
		logger.ErrorString("Redis", "IncrBy", err.Error())
		return false
	}

	return true
}

func (rds Server) Decr(key string) bool {
	if err := rds.Client.Decr(rds.Context, key).Err(); err != nil {
		logger.ErrorString("Redis", "Decr", err.Error())
		return false
	}

	return true
}

func (rds Server) DecrBy(key string, value int64) bool {
	if err := rds.Client.DecrBy(rds.Context, key, value).Err(); err != nil {
		logger.ErrorString("Redis", "DecrBy", err.Error())
		return false
	}

	return true
}
