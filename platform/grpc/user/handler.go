package user

import (
	"context"
	"go-fiber-v2/pkg/configs"
	Error "go-fiber-v2/pkg/libs/error"
	pGrpc "go-fiber-v2/pkg/libs/grpc"
	Session "go-fiber-v2/pkg/libs/session"
	"google.golang.org/grpc/connectivity"
)

type (
	UserGrpcService interface {
		TokenValidation(token string) (string, error)
	}

	userGrpc struct {
		// client to GRPC service
		session *Session.Session
	}
)

// userGrpc this type contains state of the server
type userGrpcConn struct {
	// client to GRPC service
	userClient    UserClient
	rpcConnection *pGrpc.RpcConnection
}

var rpcConnection *pGrpc.RpcConnection

func getUserClient() (*userGrpcConn, error) {
	if rpcConnection != nil && rpcConnection.Connection != nil && (rpcConnection.Connection.GetState() == connectivity.Connecting || rpcConnection.Connection.GetState() == connectivity.Ready) {
		ctx := &userGrpcConn{
			userClient:    NewUserClient(rpcConnection.Connection),
			rpcConnection: rpcConnection,
		}
		return ctx, nil

	}
	userConn, err := pGrpc.New(configs.Config.GrpcUser.Option)
	if err != nil {
		return nil, err
	}
	rpcConnection = userConn
	ctx := &userGrpcConn{
		userClient:    NewUserClient(rpcConnection.Connection),
		rpcConnection: rpcConnection,
	}
	return ctx, nil
}

// NewServerContext constructor for server context
func NewUserGrpc(session *Session.Session) UserGrpcService {
	return &userGrpc{session: session}
}

func (s *userGrpc) TokenValidation(token string) (string, error) {
	conn, err := getUserClient()
	if err != nil {
		return "", err
	}
	clientCtx, cancel := conn.rpcConnection.CreateContext(context.Background(), s.session)
	defer cancel()
	request := &RequestTokenValidation{Token: token}
	result, err := conn.userClient.TokenValidation(clientCtx, request)
	if err != nil {
		return "", Error.ParseError(err)
	}

	return result.MobilePhoneNumber, nil
}
