package dependencies

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

var ctx = context.Background()

type Cache struct {
	cfg  *Config
	conn *redis.Client
}

func NewCache(cfg *Config) *Cache {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return &Cache{
		conn: rdb,
	}
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
