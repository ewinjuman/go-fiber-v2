package cache

import (
	"github.com/redis/go-redis/v9"
	"go-fiber-v2/pkg/configs"
	"go-fiber-v2/pkg/utils"
	"golang.org/x/net/context"
)

//func init() {
//	_, err := connection()
//	if err != nil {
//		panic(fmt.Sprintf("Failed to connect redis. err: %v", err.Error()))
//	}
//}

var fiberRedisConn *redis.Client

// RedisConnection func for get connect to Redis server.
func RedisConnection() (*redis.Client, error) {
	if fiberRedisConn == nil {
		return connection()
	}
	_, err := fiberRedisConn.Ping(context.Background()).Result()
	if err != nil {
		return connection()
	}
	return fiberRedisConn, nil

}

// connection func for connect to Redis server.
func connection() (*redis.Client, error) {
	// Define Redis database number.
	dbNumber := configs.Config.Redis.Database
	// Build Redis connection URL.
	redisConnURL, err := utils.ConnectionURLBuilder("redis")
	if err != nil {
		return nil, err
	}
	options := &redis.Options{
		Addr:     redisConnURL,
		Password: configs.Config.Redis.Password,
		DB:       dbNumber,
	}
	fiberRedisConn = redis.NewClient(options)
	_, err = fiberRedisConn.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return fiberRedisConn, nil
}
