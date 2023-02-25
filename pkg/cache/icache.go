package cache

type Store interface {
	Set(key string, value string, expireTime int64) bool
	Get(key string) string
	Has(key string) bool
	Del(key string)
}
