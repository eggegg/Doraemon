package utils

import (
	"log"
	"time"

	redigo "github.com/garyburd/redigo/redis"
)

type Pool interface {
	Get() redigo.Conn
}

type Cache struct {
	MaxIdle         int
	MaxActive       int
	IdleTimeoutSecs int
	Address         string
	Auth            string
	DB              string
	Pool            *redigo.Pool
}

// NewCachePool return a new instance of the redis pool
func (cache *Cache) NewCachePool() *redigo.Pool {
		pool := &redigo.Pool{
			MaxIdle:     cache.MaxIdle,
			MaxActive:   cache.MaxActive,
			IdleTimeout: time.Second * time.Duration(cache.IdleTimeoutSecs),
			Dial: func() (redigo.Conn, error) {
				c, err := redigo.Dial("tcp", cache.Address)
				if err != nil {
					return nil, err
				}
				// if _, err = c.Do("AUTH", cache.Auth); err != nil {
				// 	c.Close()
				// 	return nil, err
				// }
				// if _, err = c.Do("SELECT", cache.DB); err != nil {
				// 	c.Close()
				// 	return nil, err
				// }
				return c, err
			},
			TestOnBorrow: func(c redigo.Conn, t time.Time) error {
				_, err := c.Do("PING")
				return err
			},
		}
		c := pool.Get() // Test connection during init
		if _, err := c.Do("PING"); err != nil {
			log.Fatal("Cannot connect to Redis: ", err)
		}
		return pool

}


func (cache *Cache) GetValue(key interface{}) (string, error) {
		conn := cache.Pool.Get()
		defer conn.Close()
		value, err := redigo.String(conn.Do("GET", key))
		return value, err
}

func (cache *Cache) SetValue(key interface{}, value interface{}) error {
	conn := cache.Pool.Get()
	defer conn.Close()
	_, err := redigo.String(conn.Do("SET", key, value))
	return err
}

func (cache *Cache) EnqueueValue(queue string, uuid string) error {
	conn := cache.Pool.Get()
	defer conn.Close()
	_, err := conn.Do("RPUSH", queue, uuid)
	return err
}
