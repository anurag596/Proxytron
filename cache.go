package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
	"os"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var redisClient *redis.Client

// init redis client
func initRedis() {
	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}

	redisClient = redis.NewClient(&redis.Options{
		Addr: redisAddr,
		DB:   0,
	})
}

// Fetch data from Redis cache
func GetFromCache(url string) (string, error) {
	result, err := redisClient.Get(ctx, url).Result()
	if err == redis.Nil {
		return "", errors.New("cache miss")
	} else if err != nil {
		return "", err
	}
	return result, nil
}

// Set data to Redis cache
func SetToCache(url, body string, expiration time.Duration) {
	err := redisClient.Set(ctx, url, body, expiration).Err()
	if err != nil {
		fmt.Printf("Failed to cache data for %s: %v\n", url, err)
	}
}

// fetch and cache the url content
func fetchAndCache(url string) (string, error) {
	// checking the cache first
	cachedResponse, err := GetFromCache(url)
	if err == nil {
		fmt.Println("Cache HIT ;)") // Cache hit
		return cachedResponse, nil
	}

	// if cache miss ,fetch from the origin
	fmt.Println("Cache MISS. Fetching from origin...")
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// store the response in cache for 2 minutes
	SetToCache(url, string(body), 2*time.Minute)
	return string(body), nil
}
