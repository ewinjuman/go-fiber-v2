package configs

import (
	GRPC "go-fiber-v2/pkg/libs/grpc"
	Rest "go-fiber-v2/pkg/libs/http"
	Logger "go-fiber-v2/pkg/libs/logger"
)

type Configuration struct {
	Apps     Apps
	Logger   Logger.Options
	Database Database
	Redis    Redis
	Ottouser Ottouser
	GrpcUser GrpcUser
}

type Apps struct {
	Name                   string
	HttpPort               int
	GrpcPort               int
	Mode                   string
	DefaultAppsId          string
	JwtSecretKey           string
	TokenExpiration        int
	JwtRefreshSecretKey    string
	RefreshTokenExpiration int
}

type Database struct {
	DbType      string `json:"dbType"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Schema      string `json:"schema"`
	Host        string `json:"host"`
	Port        int    `json:"port"`
	MaxIdleConn int    `json:"maxIdleConn"`
	MaxOpenConn int    `json:"maxOpenConn"`
	LogMode     bool   `json:"logMode"`
}

type Redis struct {
	Address  string
	Password string
	Database int
}

type Ottouser struct {
	Option Rest.Options
	Host   string
	Path   struct {
		GetUser         string `json:"getUser"`
		TokenValidation string `json:"tokenValidation"`
	}
}

type S3 struct {
	Host      string `json:"host"`
	Region    string `json:"region"`
	SecretKey string `json:"secretKey"`
	Bucket    string `json:"bucket"`
	Key       string `json:"key"`
	Timeout   int    `json:"timeout"`
}

type GrpcUser struct {
	Option GRPC.Options
}
