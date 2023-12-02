package dependencies

import (
	redispool "github.com/gomodule/redigo/redis"
	"github.com/redis/go-redis/v9"
	"time"
)

type ICache interface {
	Get(key string) (string, error)
	Set(key string, value interface{}, ttl time.Duration) error
	Pool() *redispool.Pool
}

type Cache struct {
	cfg  *Config
	conn *redis.Client
	pool *redispool.Pool
}

func NewCache(cfg *Config) *Cache {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pool := redispool.Pool{
		MaxActive: 5,
		MaxIdle:   5,
		Wait:      true,
		Dial: func() (redispool.Conn, error) {
			return redispool.Dial("tcp", cfg.RedisAddr)
		},
	}

	return &Cache{
		conn: rdb,
		pool: &pool,
	}
}

func (c *Cache) Pool() *redispool.Pool {
	return c.pool
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) error {
	return c.conn.Set(ctx, key, value, ttl).Err()
}

func (c *Cache) Get(key string) (string, error) {
	o, err := c.conn.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		return "", err
	} else {
		return o, nil
	}
}
