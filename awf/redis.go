package awf

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

//	client := redis.NewClient(&redis.Options{
//		Addr:     "127.0.0.1:9736",
//		Password: "",
//		DB:       0,
//	})
//
// ping, err := client.Ping(context.Background()).Result()

type Redis struct {
	client *redis.Client
}

func (r *Redis) setValue(key string, value interface{}, expiration time.Duration) {
	err := r.client.Set(context.Background(), key, value, expiration).Err()
	if err != nil {
		log.Println(err)
		return
	}
}

func (r *Redis) getValue(key string) interface{} {
	var data interface{}

	err := r.client.Get(context.Background(), key).Scan(&data)
	if err != nil {
		log.Println(err)
		return nil
	}

	return data
}
