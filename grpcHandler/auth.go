package grpcHandler

import (
	"io"

	"github.com/sicko7947/sicko-aio-auth/postgresql"
	auth_service "github.com/sicko7947/sicko-aio-auth/proto/auth"
)

func (s *streamService) Auth(srv auth_service.AuthStream_AuthServer) error {
	for {
		req, err := srv.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		key := req.GetKey()
		if key == "SICKOTESTTEST" {
			srv.Send(&auth_service.StreamAuthResponse{
				Code:    200,
				Message: postgresql.STATUSMAP[postgresql.OK],
			})
			continue
		}

		ipaddress := req.GetIpaddress()
		cpuId := req.GetCpuId()
		timestamp := req.GetTimestamp()

		code, _ := postgresql.Login(key, ipaddress, cpuId, timestamp)
		srv.Send(&auth_service.StreamAuthResponse{
			Code:    int64(code),
			Message: postgresql.STATUSMAP[code],
		})
	}
}

func (s *streamService) Deactivate(srv auth_service.AuthStream_DeactivateServer) error {
	for {
		req, err := srv.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		key := req.GetKey()

		code, _ := postgresql.Deactivate(key)
		srv.Send(&auth_service.StreamDeactivateResponse{
			Code:    int64(code),
			Message: postgresql.STATUSMAP[code],
		})
	}
}

func (s *streamService) Polling(srv auth_service.AuthStream_PollingServer) error {
	for {
		req, err := srv.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		key := req.GetKey()
		if key == "SICKOTESTTEST" {
			srv.Send(&auth_service.StreamPollingResponse{
				Code:    200,
				Message: postgresql.STATUSMAP[postgresql.OK],
			})
			continue
		}

		ipaddress := req.GetIpaddress()
		cpuId := req.GetCpuId()

		code, _ := postgresql.Polling(key, ipaddress, cpuId)
		srv.Send(&auth_service.StreamPollingResponse{
			Code:    int64(code),
			Message: postgresql.STATUSMAP[code],
		})
	}
}
