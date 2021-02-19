package redis

import "github.com/gomodule/redigo/redis"

var Pool *redis.Pool

func init() {
	Pool = &redis.Pool{
		MaxIdle:     16,
		MaxActive:   0,
		IdleTimeout: 300,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "192.168.101.128:6379")
		},
	}
}

func GetRedisConn() redis.Conn {
	return Pool.Get()
}
