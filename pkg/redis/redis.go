package redis

import (
	"fmt"
	"github.com/dalebao/user_control/pkg"
	"github.com/go-redis/redis"
)

var Client *redis.Client

func init() {
	dial()
}

func dial() {
	redisConfig := &setting.RedisConfig{}
	redisConfig.LoadRedis()

	Client = redis.NewClient(&redis.Options{
		Addr:     redisConfig.Host,
		Password: redisConfig.Password, // no password set
		DB:       redisConfig.Db,       // use default DB
	})

	_, err := Client.Ping().Result()
	if err != nil {
		fmt.Println("链接redis出错",err)
	}

}

