package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

var REDISPOOL *redis.Pool

func init() {
	REDISPOOL = &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		// Dial or DialContext must be set. When both are set, DialContext takes precedence over Dial.
		Dial: func() (redis.Conn, error) { return redis.Dial("tcp", "47.93.195.129:6379", redis.DialDatabase(0)) },
	}
}

func SET(str string, val string) {
	conn := REDISPOOL.Get()
	conn.Do("SET", str, val)
	conn.Close()
}
func GET(str string) string {
	conn := REDISPOOL.Get()
	defer conn.Close()
	s, err := redis.String(conn.Do("GET", str))
	if err != nil {
		fmt.Println("redis get error:", err)
		return ""
	}
	return s
}
