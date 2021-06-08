package grpcHandler

import (
	"io"

	"github.com/JasonCai686/sicko-aio-auth/postgresql"
	grpc_service "github.com/JasonCai686/sicko-aio-auth/proto/rpc"
)

func (s *streamService) RetrieveSuccess(srv grpc_service.Stream_RetrieveSuccessServer) error {
	for {
		req, err := srv.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		// get request data
		keyId := req.GetKeyId()
		successItems := postgresql.RetrieveSuccess(keyId)

		srv.Send(&grpc_service.StreamRetrieveSuccessItemsResponse{
			SuccessItems: successItems,
		})
	}
}
