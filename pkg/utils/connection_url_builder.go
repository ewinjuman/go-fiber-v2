package utils

import (
	"fmt"
	"go-fiber-v2/pkg/configs"
	"net/url"
)

// ConnectionURLBuilder func for building URL connection.
func ConnectionURLBuilder(n string) (string, error) {
	// Define URL to connection.
	var urlB string

	// Switch given names.
	switch n {
	case "postgres":
		// URL for PostgresSQL connection.
		urlB = ""
	case "mysql":
		// URL for Mysql connection.
		urlB = fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=%s",
			configs.Config.Database.Username,
			configs.Config.Database.Password,
			configs.Config.Database.Host,
			configs.Config.Database.Port,
			configs.Config.Database.Schema,
			url.QueryEscape("Asia/Jakarta"))
	case "redis":
		// URL for Redis connection.
		urlB = configs.Config.Redis.Address
	case "fiber":
		// URL for Fiber connection.
		urlB = fmt.Sprintf(
			"%s:%d",
			"0.0.0.0",
			configs.Config.Apps.HttpPort,
		)
	case "grpc":
		// URL for Redis connection.
		urlB = fmt.Sprintf(
			"tcp:%d",
			configs.Config.Apps.GrpcPort,
		)
	default:
		// Return error message.
		return "", fmt.Errorf("connection name '%v' is not supported", n)
	}

	// Return connection URL.
	return urlB, nil
}
