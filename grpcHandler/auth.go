package grpcHandler

import (
	"io"

	"github.com/JasonCai686/sicko-aio-auth/postgresql"
	grpc_service "github.com/JasonCai686/sicko-aio-auth/proto/rpc"
)

func (s *streamService) Auth(srv grpc_service.Stream_AuthServer) error {
	for {
		req, err := srv.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		key := req.GetKey()
		ipaddress := req.GetIpaddress()
		macaddress := req.GetMacaddress()
		timestamp := req.GetTimestamp()

		response := new(grpc_service.StreamAuthResponse)
		code, err := postgresql.Login(key, ipaddress, macaddress, timestamp)
		if err != nil || code != postgresql.OK {
			response.Code = int64(code)
			response.Message = postgresql.STATUSMAP[code]
			srv.Send(response)
			continue
		}
		srv.Send(response)
	}
}

func (s *streamService) Deactivate(srv grpc_service.Stream_DeactivateServer) error {
	for {
		req, err := srv.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		key := req.GetKey()

		response := new(grpc_service.StreamDeactivateResponse)
		code, err := postgresql.Deactivate(key)
		if err != nil || code != postgresql.OK {
			response.Code = int64(code)
			response.Message = postgresql.STATUSMAP[code]
			srv.Send(response)
			continue
		}
		srv.Send(response)
	}
}

func (s *streamService) Polling(srv grpc_service.Stream_PollingServer) error {
	for {
		req, err := srv.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		key := req.GetKey()

		response := new(grpc_service.StreamPollingResponse)
		code, err := postgresql.Deactivate(key)
		if err != nil || code != postgresql.OK {
			response.Code = int64(code)
			response.Message = postgresql.STATUSMAP[code]
			srv.Send(response)
			continue
		}
		srv.Send(response)
	}
}
