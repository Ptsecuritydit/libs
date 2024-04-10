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

type Context struct {
	client *redis.Client
	cnx    context.Context
}

func NewRedisContext(conf Conf, cnx context.Context) *Context {
	return &Context{
		client: redis.NewClient(&redis.Options{
			Addr:     conf.Address,
			Password: conf.Password,
			DB:       conf.DefaultDb,
		}),
		cnx: cnx,
	}
}

func (receiver *Context) Get(key string) *redis.StringCmd {
	return receiver.client.Get(receiver.cnx, key)
}

func (receiver *Context) Set(key string, value string, ttl time.Duration) *redis.StatusCmd {
	return receiver.client.Set(receiver.cnx, key, value, ttl)
}

func (receiver *Context) GetAllKeys(pattern string) *redis.Cmd {
	return receiver.client.Do(receiver.cnx, "KEYS", pattern)
}

func (receiver *Context) Subscribe(pattern string) *redis.PubSub {
	return receiver.client.PSubscribe(receiver.cnx, pattern)
}

func (receiver *Context) SetNotifyKeySpaceEvents() {
	value := `notify-keyspace-events`
	result := receiver.client.ConfigGet(receiver.cnx, value)
	if result.Err() != nil {
		log.Panic("redis config get failed", result.Err())
	}
	if result.Val()[value] == "" {
		result := receiver.client.ConfigSet(receiver.cnx, value, "KEA")
		if result.Err() != nil {
			log.Panic("redis config set failed", result.Err())
		}
	}
}
