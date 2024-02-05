package grpc

import (
	"context"
	Session "go-fiber-v2/pkg/libs/session"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"time"
)

type Options struct {
	Address string        `json:"address"`
	Timeout time.Duration `json:"timeout"`
}

type RpcConnection struct {
	options    Options
	Connection *grpc.ClientConn
}

func (rpc *RpcConnection) CreateContext(parent context.Context, session *Session.Session) (ctx context.Context, cancel context.CancelFunc) {
	ctx, cancel = context.WithTimeout(parent, rpc.options.Timeout*time.Second)
	ctx = context.WithValue(ctx, Session.AppSession, session)
	md := metadata.New(map[string]string{"Request-Id": session.ThreadID})
	ctx = metadata.NewOutgoingContext(ctx, md)
	return
}

func New(options Options) (rpc *RpcConnection, err error) {
	connection, err := grpc.Dial(options.Address, grpc.WithInsecure(), grpc.WithUnaryInterceptor(clientInterceptor))
	if err != nil {
		return nil, err
	}

	rpc = &RpcConnection{
		Connection: connection,
		options:    options,
	}
	return
}

func clientInterceptor(ctx context.Context, method string, request interface{}, response interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	timeStart := time.Now()
	session := ctx.Value(Session.AppSession).(*Session.Session)

	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		println("error meta data")
	}

	session.LogRequestGrpc(method, "GRPC", &request, md)
	err := invoker(ctx, method, request, response, cc, opts...)

	if err != nil {
		session.LogResponseGrpc(timeStart, method, "GRPC", err.Error())
		return err
	}
	session.LogResponseGrpc(timeStart, method, "GRPC", &response)
	return err
}

//=========v2

//package grpc
//
//import (
//"context"
//"time"
//
//"github.com/google/uuid"
//"google.golang.org/grpc"
//"google.golang.org/grpc/metadata"
//)
//
//type Options struct {
//	Address string        `json:"address"`
//	Timeout time.Duration `json:"timeout"`
//}
//
//type RpcConnection struct {
//	options    Options
//	Connection *grpc.ClientConn
//}
//
//func (rpc *RpcConnection) CreateContext(parent context.Context, threadID uuid.UUID) (ctx context.Context, cancel context.CancelFunc) {
//	ctx, cancel = context.WithTimeout(parent, rpc.options.Timeout)
//	md := metadata.New(map[string]string{"Request-Id": threadID.String()})
//	ctx = metadata.NewOutgoingContext(ctx, md)
//	return
//}
//
//func New(options Options) (rpc *RpcConnection, err error) {
//	connection, err := grpc.Dial(options.Address, grpc.WithInsecure(), grpc.WithUnaryInterceptor(clientInterceptor))
//	if err != nil {
//		return nil, err
//	}
//
//	rpc = &RpcConnection{
//		Connection: connection,
//		options:    options,
//	}
//	return
//}
//
//func clientInterceptor(ctx context.Context, method string, request interface{}, response interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
//	timeStart := time.Now()
//	threadID := uuid.New()
//	md := metadata.Pairs("Request-Id", threadID.String())
//	ctx = metadata.NewOutgoingContext(ctx, md)
//	err := invoker(ctx, method, request, response, cc, opts...)
//
//	if err != nil {
//		return err
//	}
//	return err
//}
