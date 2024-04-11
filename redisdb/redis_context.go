package redisdb

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

type Conf struct {
	Address   string `yaml:"address" env-default:"0.0.0.0:6379"`
	Password  string `yaml:"password"`
	DefaultDb int    `yaml:"defaultDb" env-default:"0"`
}

type RedisContext struct {
	client *redis.Client
}

func NewRedisContext(conf Conf) *RedisContext {
	return &RedisContext{
		client: redis.NewClient(&redis.Options{
			Addr:     conf.Address,
			Password: conf.Password,
			DB:       conf.DefaultDb,
		}),
	}
}

func (r *RedisContext) Get(cnx context.Context, key string) *redis.StringCmd {
	return r.client.Get(cnx, key)
}

func (r *RedisContext) Set(cnx context.Context, key string, value string, ttl time.Duration) *redis.StatusCmd {
	return r.client.Set(cnx, key, value, ttl)
}

func (r *RedisContext) GetAllKeys(cnx context.Context, pattern string) *redis.Cmd {
	return r.client.Do(cnx, "KEYS", pattern)
}

func (r *RedisContext) Subscribe(cnx context.Context, pattern string) *redis.PubSub {
	return r.client.PSubscribe(cnx, pattern)
}

func (r *RedisContext) Do(cnx context.Context, cmd string, args ...interface{}) *redis.Cmd {
	return r.client.Do(cnx, cmd, args)
}

func (r *RedisContext) SetNotifyKeySpaceEvents(cnx context.Context) {
	value := `notify-keyspace-events`
	result := r.client.ConfigGet(cnx, value)
	if result.Err() != nil {
		log.Panic("redis config get failed", result.Err())
	}
	if result.Val()[value] == "" {
		result := r.client.ConfigSet(cnx, value, "KEA")
		if result.Err() != nil {
			log.Panic("redis config set failed", result.Err())
		}
	}
}
