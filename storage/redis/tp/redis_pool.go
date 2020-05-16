package tp

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

var (
	server string = "172.31.252.134:6379" //线上
	//server string = "10.69.56.55:6379" //研发
)

//初始化连接池
func RedisPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", server,
				redis.DialConnectTimeout(2000*time.Millisecond),
				redis.DialReadTimeout(1000*time.Millisecond),
				redis.DialWriteTimeout(1000*time.Millisecond),
				//redis.DialPassword("Ksc-redis-6379")
			)
			if err != nil {
				return nil, err
			}
			return conn, err
		},
		TestOnBorrow: func(conn redis.Conn, t time.Time) error {
			_, err := conn.Do("PING")
			return err
		},
	}
}
