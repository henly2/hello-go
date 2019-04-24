package model_r

import (
	"fmt"
	"github.com/go-redis/redis"
)

var G_c *redis.Client //redis数据库的连接
func init() {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	G_c = client

}
