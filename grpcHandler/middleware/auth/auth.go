package auth

import (
	"context"
	"encoding/base64"
	"fmt"

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
	vBase64 := metautils.ExtractIncoming(ctx).Get(headerAuthorize)
	fmt.Println(vBase64)
	if vBase64 == "" {
		return nil, status.Errorf(codes.Unauthenticated, " Unauthorized")
	}

	value, _ := base64.StdEncoding.DecodeString(vBase64)
	key, err := sickocommon.RsaDecrypt(value, []byte(constants.AUTH_PRIVATE_KEY))
	if err != nil {
		fmt.Println(err)
		return nil, status.Errorf(codes.Unauthenticated, " %v", err)
	}
	fmt.Println(string(key))
	err = redis.CheckKeyExist(string(key))
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, " %v", err)
	}
	return ctx, nil
}
