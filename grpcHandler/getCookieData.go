package grpcHandler

import (
	"io"

	auth_service "github.com/sicko7947/sicko-aio-auth/proto/auth"
	"github.com/sicko7947/sicko-aio-auth/utils/redis"
)

func (s *streamService) RequestCookieData(srv auth_service.AuthStream_RequestCookieDataServer) error {
	for {
		_, err := srv.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		data := redis.GetCookieFromRedis()
		srv.Send(&auth_service.StreamGetCookieDataResponse{
			Data: data,
		})
	}
}
