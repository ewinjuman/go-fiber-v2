package grpcHandler

import (
	"context"
	"fmt"
	"go-fiber-v2/app/grpcHandler/pb"
	"go-fiber-v2/pkg/configs"
	Logger "go-fiber-v2/pkg/libs/logger"
	Session "go-fiber-v2/pkg/libs/session"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"net"
	"strconv"
)

func middleware() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, request interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		log := Logger.New(configs.Config.Logger)
		//Example for get metadata
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			println("error meta data")
		}
		values := md.Get("Request-Id")
		session := Session.New(log).
			SetRequest(&request).
			SetURL(info.FullMethod).
			SetMethod("GRPC").
			SetHeader(md)
		if len(values) > 0 {
			session.SetThreadID(values[0])
		}
		session.LogRequest(nil)
		c := context.WithValue(ctx, Session.AppSession, session)

		// TODO here, if Authentication is enable
		//errAuthenticated := status.Error(codes.Code(401), "Unauthenticated message")
		//if errAuthenticated != nil {
		//	session.LogResponse(nil, errAuthenticated.Error())
		//	return nil, errAuthenticated
		//}
		h, err := handler(c, request)
		if err != nil {
			session.LogResponse(h, err.Error())
		} else {
			session.LogResponse(h, nil)
		}
		return h, err
	}
}
func StartGrpcServer() {
	listenAddress := ":" + strconv.Itoa(configs.Config.Apps.GrpcPort)
	lis, err := net.Listen("tcp", listenAddress)
	if err != nil {
		log.Fatalf("GRPC | failed to: %v", err)
	}

	serverNew := grpc.NewServer(grpc.UnaryInterceptor(middleware()))
	pb.RegisterUserServer(serverNew, &server{})

	println(fmt.Sprintf("GRPC | server listening on %s", listenAddress))
	if err := serverNew.Serve(lis); err != nil {
		log.Fatalf("GRPC | failed to server: %v", err)
	}
}
