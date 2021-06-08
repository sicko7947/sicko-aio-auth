package auth

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"github.com/sicko7947/sicko-aio-auth/constants"
	"github.com/sicko7947/sicko-aio-auth/utils/redis"
	"github.com/sicko7947/sickocommon"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	headerAuthorize = "authorization"
)

// AuthInterceptor 认证拦截器，对以authorization为头部
func AuthInterceptor(ctx context.Context) (context.Context, error) {
	value := metautils.ExtractIncoming(ctx).Get(headerAuthorize)
	if value == "" {
		return nil, status.Errorf(codes.Unauthenticated, " Unauthorized")
	}

	key, err := sickocommon.RsaDecrypt([]byte(value), []byte(constants.AUTH_PRIVATE_KEY))
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, " %v", err)
	}

	err = redis.CheckKeyExist(string(key))
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, " %v", err)
	}
	return ctx, nil
}
