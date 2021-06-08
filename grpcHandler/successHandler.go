package grpcHandler

import (
	"io"

	"github.com/JasonCai686/sicko-aio-auth/models"
	"github.com/JasonCai686/sicko-aio-auth/postgresql"
	grpc_service "github.com/JasonCai686/sicko-aio-auth/proto/rpc"
	"github.com/JasonCai686/sicko-aio-auth/utils/webhook"
	"github.com/gogf/gf/container/gqueue"
)

var (
	successDBQueue          *gqueue.Queue
	nikeLegayQueuedWebhooks *gqueue.Queue
	nikeAcoQueuedWebhooks   *gqueue.Queue
	pacsunQueuedWebhooks    *gqueue.Queue
)

func init() {
	successDBQueue = gqueue.New()
	nikeLegayQueuedWebhooks = gqueue.New()
	nikeAcoQueuedWebhooks = gqueue.New()
	pacsunQueuedWebhooks = gqueue.New()

	go func() {
		for {
			if successDBQueue.Size() > 0 {
				if obj := successDBQueue.Pop(); obj != nil {
					item := obj.(*models.SuccessItem)
					postgresql.InsertSuccess(item)
				}
			}

			if nikeLegayQueuedWebhooks.Size() > 0 {
				if legacyItemObj := nikeLegayQueuedWebhooks.Pop(); legacyItemObj != nil {
					legacyItem := legacyItemObj.(*models.SuccessItem)
					webhook.SendLegacyNikePublicSuccess(legacyItem)
				}
			}

			if nikeAcoQueuedWebhooks.Size() > 0 {
				if acoItemObj := nikeAcoQueuedWebhooks.Pop(); acoItemObj != nil {
					acoItem := acoItemObj.(*models.SuccessItem)
					webhook.SendACONikePublicSuccess(acoItem)
				}
			}

			if pacsunQueuedWebhooks.Size() > 0 {
				if itemObj := pacsunQueuedWebhooks.Pop(); itemObj != nil {
					item := itemObj.(*models.SuccessItem)
					webhook.SendPacsunPublicSuccess(item)
				}
			}
		}
	}()
}

func (s *streamService) HandleSuccessCheckout(srv grpc_service.Stream_HandleSuccessCheckoutServer) error {
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
		setup := req.GetSetup()
		product := req.GetProduct()

		successItem := &models.SuccessItem{
			KeyId:   keyId,
			Setup:   setup,
			Product: product,
		}

		switch setup.Category {
		case "NIKE":
			switch product.MerchGroup {
			case "XP", "XA", "MX":
				nikeLegayQueuedWebhooks.Push(successItem)
			default:
				nikeAcoQueuedWebhooks.Push(successItem)
			}
		}
		successDBQueue.Push(successItem)

		srv.Send(&grpc_service.StreamHandleSuccessCheckoutResponse{
			Success: true,
			Errors:  nil,
		})
	}
}
