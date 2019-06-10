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
	fmt.Println(redisConfig)

	Client = redis.NewClient(&redis.Options{
		Addr:     redisConfig.Host,
		Password: redisConfig.Password, // no password set
		DB:       redisConfig.Db,       // use default DB
	})

	pong, err := Client.Ping().Result()
	fmt.Println(pong, err)
}

func Test() {
	Client.Set("test", "112", 0)
}
