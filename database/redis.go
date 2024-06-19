package database

import (
	"context"
	"crypto/tls"
	"log"
	"net/url"

	"github.com/go-redis/redis/v8"

	"github.com/souvik150/BattleQuiz-Backend/config"
)

var RedisClient *redis.Client
var Ctx = context.Background()

func ConnectRedis() error {
	redisURL := config.GetRedisConfig()
	urlParsed, err := url.Parse(redisURL)
	if err != nil {
		log.Fatal("Failed to parse Redis URL:", err)
		return err
	}

	password, _ := urlParsed.User.Password()
	addr := urlParsed.Host

	RedisClient = redis.NewClient(&redis.Options{
		Addr:      addr,
		Password:  password,
		DB:        0,
		TLSConfig: &tls.Config{},
	})

	_, err = RedisClient.Ping(Ctx).Result()
	if err != nil {
		log.Fatal("Failed to connect to Redis:", err)
		return err
	}

	return nil
}
