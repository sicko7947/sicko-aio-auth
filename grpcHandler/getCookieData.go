package grpcHandler

import (
	auth_service "github.com/sicko7947/sicko-aio-auth/proto/auth"
	"github.com/sicko7947/sicko-aio-auth/utils/redis"
)

func (s *streamService) RequestCookieData(srv auth_service.Stream_RequestCookieDataServer) error {
	for {
		_, err := srv.Recv()
		if err != nil {
			return err
		}

		data := redis.GetCookieFromRedis()
		if len(data) > 0 {
			srv.Send(&auth_service.StreamGetCookieDataResponse{
				Data:   data,
				Errors: nil,
			})
		} else {
			srv.Send(&auth_service.StreamGetCookieDataResponse{
				Errors: &auth_service.Errors{
					Code:    400,
					Message: "Error Getting Sicko Cookies",
				},
			})
			continue
		}
	}
}
