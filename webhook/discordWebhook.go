package discordwebhook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/sicko7947/sicko-aio-auth/models"
	"github.com/sicko7947/sicko-aio-auth/utils/psychoclient"
)

var (
	session psychoclient.Session

	webhooks_MX = []string{
		"https://discord.com/api/webhooks/xxx/xxx",
	}
	webhooks_XP = []string{
		"https://discord.com/api/webhooks/xxx/xxx",
	}
	webhooks_XA = []string{
		"https://discord.com/api/webhooks/xxx/xxx",
	}
	webhooks_EU = []string{
		"https://discord.com/api/webhooks/xxx/xxx",
	}
	webhooks_US = []string{
		"https://discord.com/api/webhooks/xxx/xxx",
	}
	webhooks_CN = []string{
		"https://discord.com/api/webhooks/xxx/xxx",
	}
	webhooks_JP = []string{
		"https://discord.com/api/webhooks/xxx/xxx",
	}
	webhooks_SSENSE = []string{
		"https://discord.com/api/webhooks/xxx/xxx",
	}
)

func init() {
	session, _ = psychoclient.NewSession(&psychoclient.SessionBuilder{
		UseDefaultClient: true,
	})
}

func send(webhookUrl string, payload []byte) {
	reqId, _ := session.BuildRequest(&psychoclient.RequestBuilder{
		Endpoint: webhookUrl,
		Method:   "POST",
		Headers: map[string]string{
			"content-type": "application/json; charset=UTF-8",
			"accept":       "application/json; charset=UTF-8, application/json",
		},
		Payload: bytes.NewBuffer(payload),
	})

	sleepTime := rand.Intn(2000-1000) + 1000
	ticker := time.NewTicker(time.Duration(sleepTime * int(time.Millisecond)))
	for {
		res, _, err := session.Do(reqId, false)
		if err != nil {
			return
		}
		switch res.StatusCode {
		case 204:
			session.RemoveRequest(reqId)
			return
		case 401:
			return
		case 429:
			<-ticker.C
		}
	}
}

func SendLegacyNikePublicSuccess(successItem *models.SuccessItem) {
	faker := gofakeit.New(0)

	data, _ := json.Marshal(&models.WebhookBuilder{
		Embeds: []*models.Embed{
			{
				Color:       "65419",
				Title:       successItem.Product.ProductName,
				Description: successItem.Product.ProductDescription,
				Fields: []*models.EmbedField{
					{
						Name:   "Category",
						Value:  successItem.Setup.Category,
						Inline: true,
					},
					{
						Name:   "Region",
						Value:  successItem.Setup.Region,
						Inline: true,
					},
					{
						Name:   "\u200b",
						Value:  "\u200b",
						Inline: true,
					},
					{
						Name:   "Product SKU",
						Value:  successItem.Product.ProductSku,
						Inline: true,
					},
					{
						Name:   "Size",
						Value:  successItem.Product.Size,
						Inline: true,
					},
					{
						Name:   "Quantity",
						Value:  fmt.Sprint(successItem.Product.Quantity),
						Inline: true,
					},
					{
						Name: "Price",
						Value: func() (price string) {
							price = successItem.Product.Price
							if len(price) == 0 {
								price = "N/A"
							}
							return price
						}(),
						Inline: true,
					},
					{
						Name:   "Time",
						Value:  successItem.Setup.Timestamp,
						Inline: false,
					},
					{
						Name:   "Task Type",
						Value:  successItem.Setup.TaskType,
						Inline: true,
					},
				},
				Thumbnail: &models.EmbedThumbnail{
					URL:    "https://secure-images.nike.com/is/image/DotCom/" + strings.ReplaceAll(successItem.Product.ProductSku, "-", "_"),
					Width:  400,
					Height: 400,
				},
				Footer: &models.EmbedFooter{
					Text:         fmt.Sprintf("Sicko AIO - 2.0 [%s]", time.Now().Format(time.RFC3339Nano)),
					IconURL:      "https://pbs.twimg.com/profile_images/1122681028210905088/2cZIhvv-_400x400.png",
					ProxyIconURL: "https://images-ext-1.discordapp.net/external/p8C-Btf5KSrbr1YkqPvgl980BPQ8PDLyJ4Le1paGn1M/https/pbs.twimg.com/profile_images/1122681028210905088/2cZIhvv-_400x400.png",
				},
			},
		},
	})

	var webhookUrl string
	switch successItem.Product.MerchGroup {
	case "XA":
		webhookUrl = faker.RandomString(webhooks_XA)
	case "XP":
		webhookUrl = faker.RandomString(webhooks_XP)
	case "MX":
		webhookUrl = faker.RandomString(webhooks_MX)
	}
	go send(webhookUrl, data)
}

func SendACONikePublicSuccess(successItem *models.SuccessItem) {
	faker := gofakeit.New(0)

	var useGiftCard, useDiscount, guest bool
	if len(successItem.Product.GiftCards) > 0 {
		useGiftCard = true
	}
	if len(successItem.Product.Account) == 0 {
		guest = true
	}
	if len(successItem.Product.DiscountCode) > 0 {
		useGiftCard = true
	}

	data, _ := json.Marshal(&models.WebhookBuilder{
		Embeds: []*models.Embed{
			{
				Color:       "65419",
				Title:       successItem.Product.ProductName,
				Description: successItem.Product.ProductDescription,
				Fields: []*models.EmbedField{
					{
						Name:   "Category",
						Value:  successItem.Setup.Category,
						Inline: true,
					},
					{
						Name:   "Region",
						Value:  successItem.Setup.Region,
						Inline: true,
					},
					{
						Name:   "\u200b",
						Value:  "\u200b",
						Inline: true,
					},
					{
						Name:   "Product SKU",
						Value:  successItem.Product.ProductSku,
						Inline: true,
					},
					{
						Name:   "Size",
						Value:  successItem.Product.Size,
						Inline: true,
					},
					{
						Name:   "Quantity",
						Value:  fmt.Sprint(successItem.Product.Quantity),
						Inline: true,
					},
					{
						Name: "Price",
						Value: func() (price string) {
							price = successItem.Product.Price
							if len(price) == 0 {
								price = "N/A"
							}
							return price
						}(),
						Inline: false,
					},
					{
						Name:   "Guest",
						Value:  fmt.Sprint(guest),
						Inline: true,
					},
					{
						Name:   "GiftCard",
						Value:  fmt.Sprint(useGiftCard),
						Inline: true,
					},
					{
						Name:   "\u200b",
						Value:  "\u200b",
						Inline: true,
					},
					{
						Name:   "Discount",
						Value:  fmt.Sprint(useDiscount),
						Inline: true,
					},
					{
						Name:   "Psycho Cookie",
						Value:  fmt.Sprint(successItem.Setup.UsePsychoCookie),
						Inline: true,
					},
					{
						Name:   "Monitor Mode",
						Value:  successItem.Setup.MonitorMode,
						Inline: false,
					},
					{
						Name:   "Time",
						Value:  successItem.Setup.Timestamp,
						Inline: true,
					},
					{
						Name:   "Task Type",
						Value:  successItem.Setup.TaskType,
						Inline: false,
					},
				},
				Thumbnail: &models.EmbedThumbnail{
					URL:    "https://secure-images.nike.com/is/image/DotCom/" + strings.ReplaceAll(successItem.Product.ProductSku, "-", "_"),
					Width:  400,
					Height: 400,
				},
				Footer: &models.EmbedFooter{
					Text:         fmt.Sprintf("Sicko AIO - 2.0 [%s]", time.Now().Format("2006-01-02T15:04:05.000Z")),
					IconURL:      "https://pbs.twimg.com/profile_images/1122681028210905088/2cZIhvv-_400x400.png",
					ProxyIconURL: "https://images-ext-1.discordapp.net/external/p8C-Btf5KSrbr1YkqPvgl980BPQ8PDLyJ4Le1paGn1M/https/pbs.twimg.com/profile_images/1122681028210905088/2cZIhvv-_400x400.png",
				},
			},
		},
	})

	var webhookUrl string
	switch successItem.Product.MerchGroup {
	case "EU":
		webhookUrl = faker.RandomString(webhooks_EU)
	case "US":
		webhookUrl = faker.RandomString(webhooks_US)
	case "CN":
		webhookUrl = faker.RandomString(webhooks_CN)
	case "JP":
		webhookUrl = faker.RandomString(webhooks_JP)
	}
	go send(webhookUrl, data)
}

func SendPacsunPublicSuccess(successItem *models.SuccessItem) {
	faker := gofakeit.New(0)

	data, _ := json.Marshal(&models.WebhookBuilder{
		Embeds: []*models.Embed{
			{
				Color:       "65419",
				Title:       successItem.Product.ProductName,
				Description: successItem.Product.ProductDescription,
				Fields: []*models.EmbedField{
					{
						Name:   "Category",
						Value:  successItem.Setup.Category,
						Inline: true,
					},
					{
						Name:   "Region",
						Value:  successItem.Setup.Region,
						Inline: true,
					},
					{
						Name:   "\u200b",
						Value:  "\u200b",
						Inline: true,
					},
					{
						Name:   "Product SKU",
						Value:  successItem.Product.ProductSku,
						Inline: true,
					},
					{
						Name:   "Size",
						Value:  successItem.Product.Size,
						Inline: true,
					},
					{
						Name:   "Quantity",
						Value:  fmt.Sprint(successItem.Product.Quantity),
						Inline: true,
					},
					{
						Name: "Price",
						Value: func() (price string) {
							price = successItem.Product.Price
							if len(price) == 0 {
								price = "N/A"
							}
							return price
						}(),
						Inline: false,
					},
					{
						Name:   "Time",
						Value:  successItem.Setup.Timestamp,
						Inline: false,
					},
					{
						Name:   "Task Type",
						Value:  successItem.Setup.TaskType,
						Inline: false,
					},
				},
				Thumbnail: &models.EmbedThumbnail{
					URL:    "",
					Width:  400,
					Height: 400,
				},
				Footer: &models.EmbedFooter{
					Text:         fmt.Sprintf("Sicko AIO - 2.0 [%s]", time.Now().Format("2006-01-02T15:04:05.000Z")),
					IconURL:      "https://pbs.twimg.com/profile_images/1122681028210905088/2cZIhvv-_400x400.png",
					ProxyIconURL: "https://images-ext-1.discordapp.net/external/p8C-Btf5KSrbr1YkqPvgl980BPQ8PDLyJ4Le1paGn1M/https/pbs.twimg.com/profile_images/1122681028210905088/2cZIhvv-_400x400.png",
				},
			},
		},
	})

	var webhookUrl string
	switch successItem.Product.MerchGroup {
	case "EU":
		webhookUrl = faker.RandomString(webhooks_EU)
	case "US":
		webhookUrl = faker.RandomString(webhooks_US)
	case "CN":
		webhookUrl = faker.RandomString(webhooks_CN)
	case "JP":
		webhookUrl = faker.RandomString(webhooks_JP)
	}
	go send(webhookUrl, data)
}

func SendSsensePublicSuccess(successItem *models.SuccessItem) {
	faker := gofakeit.New(0)

	data, _ := json.Marshal(&models.WebhookBuilder{
		Embeds: []*models.Embed{
			{
				Color:       "65419",
				Title:       successItem.Product.ProductName,
				Description: successItem.Product.ProductDescription,
				Fields: []*models.EmbedField{
					{
						Name:   "Category",
						Value:  successItem.Setup.Category,
						Inline: true,
					},
					{
						Name:   "Region",
						Value:  successItem.Setup.Region,
						Inline: true,
					},
					{
						Name:   "\u200b",
						Value:  "\u200b",
						Inline: true,
					},
					{
						Name:   "Product SKU",
						Value:  successItem.Product.ProductSku,
						Inline: true,
					},
					{
						Name:   "Size",
						Value:  successItem.Product.Size,
						Inline: true,
					},
					{
						Name:   "Quantity",
						Value:  fmt.Sprint(successItem.Product.Quantity),
						Inline: true,
					},
					{
						Name: "Price",
						Value: func() (price string) {
							price = successItem.Product.Price
							if len(price) == 0 {
								price = "N/A"
							}
							return price
						}(),
						Inline: false,
					},
					{
						Name:   "Time",
						Value:  successItem.Setup.Timestamp,
						Inline: false,
					},
					{
						Name:   "Task Type",
						Value:  successItem.Setup.TaskType,
						Inline: false,
					},
				},
				Thumbnail: &models.EmbedThumbnail{
					URL:    successItem.Product.ImageUrl,
					Width:  400,
					Height: 400,
				},
				Footer: &models.EmbedFooter{
					Text:         fmt.Sprintf("Sicko AIO - 2.0 [%s]", time.Now().Format("2006-01-02T15:04:05.000Z")),
					IconURL:      "https://pbs.twimg.com/profile_images/1122681028210905088/2cZIhvv-_400x400.png",
					ProxyIconURL: "https://images-ext-1.discordapp.net/external/p8C-Btf5KSrbr1YkqPvgl980BPQ8PDLyJ4Le1paGn1M/https/pbs.twimg.com/profile_images/1122681028210905088/2cZIhvv-_400x400.png",
				},
			},
		},
	})

	webhookUrl := faker.RandomString(webhooks_SSENSE)
	go send(webhookUrl, data)
}
