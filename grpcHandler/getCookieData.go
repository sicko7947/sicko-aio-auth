package grpcHandler

import (
	grpc_service "github.com/JasonCai686/sicko-aio-auth/proto/rpc"
)

func (s *streamService) RequestCookieData(srv grpc_service.Stream_RequestCookieDataServer) error {
	for {
		_, err := srv.Recv()
		if err != nil {
			return err
		}

		data := utils.GetCookie2FromRedis()
		if len(data) > 0 {
			srv.Send(&grpc_service.StreamGetCookieDataResponse{
				Data:   data,
				Errors: nil,
			})
		} else {
			srv.Send(&grpc_service.StreamGetCookieDataResponse{
				Errors: &grpc_service.Errors{
					Code:    400,
					Message: "Error Getting Sicko Cookies",
				},
			})
			continue
		}
	}
}
