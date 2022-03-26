package Redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

func Set(key string, val string, expiration time.Duration) error {
	return redisDB.Set(key, val, expiration).Err()
}

// key
func Get(key string) (string, error) {
	return redisDB.Get(key).Result()
}

func Zset(key string, val []redis.Z) error {
	_, err := redisDB.ZAdd(key, val...).Result()
	if err != nil {
		fmt.Printf("zadd failed, err:%v\n", err)
		return err
	}
	return nil
}

//Queue
func LPush(key string, val string) {

	BaseCmd := redisDB.LPush(key, val)

	fmt.Println(key, "has", BaseCmd.Val(), "value")

	return
}

//getList
func LRange(key string, loc_s int64, loc_e int64) []string {

	cmd := redisDB.LRange(key, loc_s, loc_e)

	return cmd.Val()

}
