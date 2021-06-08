package grpcHandler

import (
	"log"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/shimingyah/pool"
	"github.com/sicko7947/sicko-aio-auth/grpcHandler/middleware/auth"
	"github.com/sicko7947/sicko-aio-auth/grpcHandler/middleware/cred"
	"github.com/sicko7947/sicko-aio-auth/grpcHandler/middleware/recovery"
	"github.com/sicko7947/sicko-aio-auth/grpcHandler/middleware/zap"
	auth_service "github.com/sicko7947/sicko-aio-auth/proto/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

var (
	AuthServer *grpc.Server
)

type streamService struct{}

func StargrpcServer(port string) {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	// Create a new grpc server
	AuthServer = grpc.NewServer(
		cred.TLSInterceptor(),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_validator.StreamServerInterceptor(),
			grpc_zap.StreamServerInterceptor(zap.ZapInterceptor()),
			grpc_auth.StreamServerInterceptor(auth.AuthInterceptor),
			grpc_recovery.StreamServerInterceptor(recovery.RecoveryInterceptor()),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_validator.UnaryServerInterceptor(),
			grpc_zap.UnaryServerInterceptor(zap.ZapInterceptor()),
			grpc_auth.UnaryServerInterceptor(auth.AuthInterceptor),
			grpc_recovery.UnaryServerInterceptor(recovery.RecoveryInterceptor()),
		)),

		// sete grpc connection pool configuration
		grpc.InitialWindowSize(pool.InitialWindowSize),
		grpc.InitialConnWindowSize(pool.InitialConnWindowSize),
		grpc.MaxSendMsgSize(pool.MaxSendMsgSize),
		grpc.MaxRecvMsgSize(pool.MaxRecvMsgSize),
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			PermitWithoutStream: true,
		}),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			Time:    pool.KeepAliveTime,
			Timeout: pool.KeepAliveTimeout,
		}),
	)
	auth_service.RegisterStreamServer(AuthServer, &streamService{})
	log.Println(port + " HTTP.Listing whth TLS and token...")
	err = AuthServer.Serve(listener)
	if err != nil {
		log.Fatalf("grpcServer.Serve err: %v", err)
	}
}
