package ws

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient(redisHost string) *RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr: redisHost,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	return &RedisClient{client: rdb}
}

func (r *RedisClient) SaveMessage(msg *Message) {
	err := r.client.RPush(ctx, "messages", msg.Content).Err()
	if err != nil {
		log.Printf("Error saving message to Redis: %v", err)
	}
}

func (r *RedisClient) GetMessages() ([]string, error) {
	messages, err := r.client.LRange(ctx, "messages", 0, -1).Result()
	if err != nil {
		log.Printf("Error retrieving messages from Redis: %v", err)
		return nil, err
	}
	return messages, nil
}
