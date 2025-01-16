package awf

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

func getRedisC() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:9736",
		Password: "",
		DB:       0,
	})

	ping, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Println(err)
		return nil
	}

	log.Println(ping)
	return client
}

func setValue(key string, value interface{}, expiration time.Duration) {
	client := getRedisC()
	err := client.Set(context.Background(), key, value, expiration).Err()
	if err != nil {
		log.Println(err)
		return
	}
}

func getValue(key string) interface{} {
	var data interface{}

	client := getRedisC()
	err := client.Get(context.Background(), key).Scan(&data)
	if err != nil {
		log.Println(err)
		return nil
	}

	return data
}
