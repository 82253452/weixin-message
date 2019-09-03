package redis

import (
	"awesomeProject/src/tool"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"strconv"
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

func LRANGE() []string {
	conn := REDISPOOL.Get()
	defer conn.Close()
	size, _ := strconv.Atoi(tool.Config.Redis.Size)
	list, err := redis.Strings(conn.Do("LRANGE", tool.Config.Redis.ParseKey, 0, size-1))
	if err != nil {
		fmt.Println("redis get error:", err)
		return nil
	}
	return list
}

func LTRIM() {
	conn := REDISPOOL.Get()
	size, _ := strconv.Atoi(tool.Config.Redis.Size)
	r, err := conn.Do("LTRIM", tool.Config.Redis.ParseKey, size-1, -1)
	if err != nil {
		fmt.Println("redis get error:", err)
	}
	fmt.Println("LTRIM:", r)
	conn.Close()
}

func LREM() []string {
	conn := REDISPOOL.Get()
	defer conn.Close()
	list, err := redis.Strings(conn.Do("LREM", tool.Config.Redis.ParseKey, 0, tool.Config.Redis.Size))
	if err != nil {
		fmt.Println("redis get error:", err)
		return nil
	}
	return list
}
