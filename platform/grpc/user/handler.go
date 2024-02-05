package user

import (
	"context"
	"go-fiber-v2/pkg/configs"
	Error "go-fiber-v2/pkg/libs/error"
	pGrpc "go-fiber-v2/pkg/libs/grpc"
	Session "go-fiber-v2/pkg/libs/session"
	"google.golang.org/grpc/connectivity"
)

// serverContext this type contains state of the server
type serverContext struct {
	// client to GRPC service
	userClient    UserClient
	rpcConnection *pGrpc.RpcConnection
}

var rpcConnection *pGrpc.RpcConnection

// NewServerContext constructor for server context
func NewServerContext() (*serverContext, error) {
	if rpcConnection != nil && rpcConnection.Connection != nil && (rpcConnection.Connection.GetState() == connectivity.Connecting || rpcConnection.Connection.GetState() == connectivity.Ready) {
		ctx := &serverContext{
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
	ctx := &serverContext{
		userClient:    NewUserClient(userConn.Connection),
		rpcConnection: userConn,
	}
	return ctx, nil
}

func (s *serverContext) TokenValidation(session *Session.Session, token string) (string, error) {
	clientCtx, cancel := s.rpcConnection.CreateContext(context.Background(), session)
	defer cancel()
	request := &RequestTokenValidation{Token: token}
	result, err := s.userClient.TokenValidation(clientCtx, request)
	if err != nil {
		return "", Error.ParseError(err)
	}

	return result.MobilePhoneNumber, nil
}

func (s *serverContext) GetUser(session *Session.Session, mobilePhoneNumber string) (*ResponseGetUser, error) {
	clientCtx, cancel := s.rpcConnection.CreateContext(context.Background(), session)
	defer cancel()
	request := &RequestGetUser{MobilePhoneNumber: mobilePhoneNumber}
	result, err := s.userClient.GetUser(clientCtx, request)
	if err != nil {
		return nil, Error.ParseError(err)
	}
	return result, nil
}

func (s *serverContext) GetCivilRegistry(session *Session.Session, mobilePhoneNumber string) (*ResponseGetCivilRegistry, error) {
	clientCtx, cancel := s.rpcConnection.CreateContext(context.Background(), session)
	defer cancel()
	request := &RequestGetCivilRegistry{MobilePhoneNumber: mobilePhoneNumber}
	result, err := s.userClient.GetCivilRegistry(clientCtx, request)
	if err != nil {
		return nil, Error.ParseError(err)
	}
	return result, nil
}
