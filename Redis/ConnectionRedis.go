package Redis

import (
	"fmt"
	"github.com/go-redis/redis"
)

var redisDB *redis.Client

const Address = "localhost"
const Port = "6379"
const Password = ""
const DB = 0

func InitRedis() {

	redisDB = redis.NewClient(&redis.Options{
		Addr:     Address + ":" + Port,
		Password: Password,
		DB:       DB,
	})

	_, err := redisDB.Ping().Result()

	if err != nil {
		fmt.Println(err)
	}

}
