package cache

import (
	"github.com/gofiber/storage/redis/v3"
	"go-fiber-v2/pkg/configs"
	"go-fiber-v2/pkg/utils"
	"golang.org/x/net/context"
	"runtime"
)

var fiberRedisConn *redis.Storage

func init() {
	//_, err := connection()
	//if err != nil {
	//	panic(fmt.Sprintf("Failed to connect redis. err: %v", err.Error()))
	//}
	//rTTL := time.Duration(0) * time.Minute
	//rTTL, _ := time.ParseDuration("1fg")
	//err = fiberRedisConn.Set("KEY1", []byte("ini data"), rTTL)
	//if err != nil {
	//	println(err.Error())
	//}
}

// RedisConnection func for get connect to Redis server.
func RedisConnection() (*redis.Storage, error) {
	if fiberRedisConn == nil {
		return connection()
	}
	_, err := fiberRedisConn.Conn().Ping(context.Background()).Result()
	if err != nil {
		return connection()
	}
	return fiberRedisConn, nil

}

// connection func for connect to Redis server.
func connection() (*redis.Storage, error) {
	// Define Redis database number.
	dbNumber := configs.Config.Redis.Database
	// Build Redis connection URL.
	redisConnURL, err := utils.ConnectionURLBuilder("redis")
	if err != nil {
		return nil, err
	}
	store := redis.New(redis.Config{
		Addrs:    []string{redisConnURL},
		Username: "",
		Password: "",
		Database: dbNumber,
		PoolSize: 10 * runtime.GOMAXPROCS(0),
	})

	fiberRedisConn = store
	_, err = fiberRedisConn.Conn().Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return fiberRedisConn, nil
}
