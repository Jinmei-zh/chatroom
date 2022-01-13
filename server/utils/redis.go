package utils

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

func InitPool(address string, MaxIdle, MaxActive int, IdleTimeout time.Duration) (pool *redis.Pool) {
	pool = &redis.Pool{
		MaxIdle:     MaxIdle,
		MaxActive:   MaxActive,
		IdleTimeout: IdleTimeout,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", address)
		},
	}
	return
}
