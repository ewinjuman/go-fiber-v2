package grpc

//import (
//"context"
//Config "go-otto-users/utils/config"
//Logger "go-otto-users/utils/logger"
//Session "go-otto-users/utils/session"
//"google.golang.org/grpc"
//"google.golang.org/grpc/metadata"
//"net"
//"strconv"
//"time"
//)
//
//type server struct{}
//
//func middleware() grpc.UnaryServerInterceptor {
//	return func(ctx context.Context, request interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
//		log := Logger.New(Config.Config.Logger)
//		//Example for get metadata
//		md, ok := metadata.FromIncomingContext(ctx)
//		if !ok {
//			println("error meta data")
//		}
//		values := md.Get("Request-Id")
//		sessionID := strconv.Itoa(int(time.Now().UnixNano() / int64(time.Millisecond)))
//		if len(values) > 0 {
//			sessionID = values[0]
//		}
//		session := Session.New(log).
//			SetThreadID(sessionID).
//			SetRequest(&request).
//			SetURL(info.FullMethod).
//			SetMethod("GRPC").
//			SetHeader(md)
//		session.LogRequest(nil)
//		c := context.WithValue(ctx, Session.AppSession, session)
//
//		// TODO here, if Authentication is enable
//		//errAuthenticated := status.Error(codes.Code(401), "Unauthenticated message")
//		//if errAuthenticated != nil {
//		//	session.LogResponse(nil, errAuthenticated.Error())
//		//	return nil, errAuthenticated
//		//}
//		h, err := handler(c, request)
//		if err != nil {
//			session.LogResponse(h, err.Error())
//		} else {
//			session.LogResponse(h, nil)
//		}
//		return h, err
//	}
//}
//func StartGrpcServer() {
//	listenAddress := ":" + strconv.Itoa(Config.Config.Apps.GrpcPort) // TODO: create config grpc port
//	lis, err := net.Listen("tcp", listenAddress)
//	if err != nil {
//		log.Fatalf("GRPC | failed to: %v", err)
//	}
//
//	serverNew := grpc.NewServer(grpc.UnaryInterceptor(middleware()))
//	RegisterUserServer(serverNew, &server{})
//
//	println(fmt.Sprintf("GRPC | server listening on %s", listenAddress))
//	if err := serverNew.Serve(lis); err != nil {
//		log.Fatalf("GRPC | failed to server: %v", err)
//	}
//}
//
//func UnaryClientInterceptor() grpc.UnaryClientInterceptor {
//	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
//		err := invoker(ctx, method, req, reply, cc, opts...)
//		return err
//	}
//}
