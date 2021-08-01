package grpcHandler

import (
	"io"
	"time"

	"github.com/gogf/gf/container/gqueue"
	"github.com/sicko7947/sicko-aio-auth/models"
	"github.com/sicko7947/sicko-aio-auth/postgresql"
	auth_service "github.com/sicko7947/sicko-aio-auth/proto/auth"
	discordwebhook "github.com/sicko7947/sicko-aio-auth/webhook"
)

var (
	successDBQueue          *gqueue.Queue
	nikeLegayQueuedWebhooks *gqueue.Queue
	nikeAcoQueuedWebhooks   *gqueue.Queue
	pacsunQueuedWebhooks    *gqueue.Queue
	ssenseQueuedWebhooks    *gqueue.Queue
)

func init() {
	successDBQueue = gqueue.New()
	nikeLegayQueuedWebhooks = gqueue.New()
	nikeAcoQueuedWebhooks = gqueue.New()
	pacsunQueuedWebhooks = gqueue.New()
	ssenseQueuedWebhooks = gqueue.New()

	go func() {
		ticker := time.NewTicker(100 * time.Millisecond)
		for {
			if successDBQueue.Size() > 0 {
				if obj := successDBQueue.Pop(); obj != nil {
					item := obj.(*models.SuccessItem)
					go postgresql.InsertSuccess(item)
				}
			}

			if nikeLegayQueuedWebhooks.Size() > 0 {
				if legacyItemObj := nikeLegayQueuedWebhooks.Pop(); legacyItemObj != nil {
					legacyItem := legacyItemObj.(*models.SuccessItem)
					go discordwebhook.SendLegacyNikePublicSuccess(legacyItem)
				}
			}

			if nikeAcoQueuedWebhooks.Size() > 0 {
				if acoItemObj := nikeAcoQueuedWebhooks.Pop(); acoItemObj != nil {
					acoItem := acoItemObj.(*models.SuccessItem)
					go discordwebhook.SendACONikePublicSuccess(acoItem)
				}
			}

			if pacsunQueuedWebhooks.Size() > 0 {
				if itemObj := pacsunQueuedWebhooks.Pop(); itemObj != nil {
					item := itemObj.(*models.SuccessItem)
					go discordwebhook.SendPacsunPublicSuccess(item)
				}
			}

			if ssenseQueuedWebhooks.Size() > 0 {
				if itemObj := ssenseQueuedWebhooks.Pop(); itemObj != nil {
					item := itemObj.(*models.SuccessItem)
					go discordwebhook.SendSsensePublicSuccess(item)
				}
			}
			<-ticker.C
		}
	}()
}

func (s *streamService) HandleSuccessCheckout(srv auth_service.AuthStream_HandleSuccessCheckoutServer) error {
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
		case "SSENSE":
			ssenseQueuedWebhooks.Push(successItem)
		}
		successDBQueue.Push(successItem)

		srv.Send(&auth_service.StreamHandleSuccessCheckoutResponse{
			Success: true,
		})
	}
}
