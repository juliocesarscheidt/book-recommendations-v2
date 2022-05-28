package adapter

import (
	"fmt"
	"time"
	"errors"
	"context"

	redis "github.com/go-redis/redis/v8"
)

type RedisClient struct {
	Client *redis.Client
}

func GetRedisClient(redisConnString string) RedisClient {
	client, err := GetClient(redisConnString)
	if err != nil {
		fmt.Println("[ERROR]", err)
	}
	return RedisClient{
		Client: client,
	}
}

func GetClient(redisConnString string) (*redis.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	redisOpts := &redis.Options{
		Addr:     redisConnString,
		Password: "", // default no password
		DB:       0,
	}
	redisClient := redis.NewClient(redisOpts)
	if redisClient == nil {
		return nil, errors.New("Failed to connect to Redis endpoint")
	}
	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		fmt.Println("[ERROR]", err)
		return nil, errors.New("Failed to connect to Redis endpoint")
	}

	return redisClient, nil
}

func (redisClient RedisClient) Get(key string) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	result, err := redisClient.Client.Get(ctx, string(key)).Result()
	if err == redis.Nil {
		return nil, err
	} else if err != nil {
		fmt.Println("[ERROR]", err)
		return nil, err
	}

	return result, nil
}

func (redisClient RedisClient) Set(key string, value interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	expiryTime := 60*60*time.Second // 1 hour
	if err := redisClient.Client.Set(ctx, key, fmt.Sprintf("%v", value), expiryTime).Err(); err != nil {
		fmt.Println("[ERROR]", err)
		return err
	}

	return nil
}

func (redisClient RedisClient) Del(key string) (error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := redisClient.Client.Del(ctx, key).Err(); err != nil {
		fmt.Println("[ERROR]", err)
		return err
	}

	return nil
}
