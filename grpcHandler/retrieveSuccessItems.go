package grpcHandler

import (
	"io"

	"github.com/sicko7947/sicko-aio-auth/postgresql"
	auth_service "github.com/sicko7947/sicko-aio-auth/proto/auth"
)

func (s *streamService) RetrieveSuccess(srv auth_service.AuthStream_RetrieveSuccessServer) error {
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

		srv.Send(&auth_service.StreamRetrieveSuccessItemsResponse{
			SuccessItems: successItems,
		})
	}
}
