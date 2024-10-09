package initializers

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client
var Ctx = context.Background()

func ConnectToRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping(Ctx).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pong)
	RedisClient = client
}
