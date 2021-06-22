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
		"https://discord.com/api/webhooks/838240551912341504/--H3i475kCzpxPQ9M2Vp43G3yDjEyGQNhzwCrX29UcCt9G1ElFoc-5Z_PeuLs9JVq3lY",
		"https://discord.com/api/webhooks/838240565937831966/l9sMvnjvjzPFvYQ9YzXvFvF9VIwa6GBFMmmixgYSrt-MVOfyVOTab_kYNEA5mwtaoR7H",
		"https://discord.com/api/webhooks/838240584845492255/2iK_lHjRG6h7XlaEBNz_pspc-U-J20CSF8sqRc38hF6t8INxq7bwW94DL7bvf2vz6-TR",
		"https://discord.com/api/webhooks/838240805679398943/MGKTjHnCMSbJV2AF6rZpyowdFSToPN1WTMh63-q77j1O4x_mQypcjbIh4vmijhYX5oww",
		"https://discord.com/api/webhooks/838240829171826690/H1gt3QV7DZOSTIJw3oRwr8fhvWoyexFPlZNGAXISGY9KLBTb4qNxRL3bsdGkaCnbp8B2",
		"https://discord.com/api/webhooks/838240846578974741/mS_3G9c5QQpsCA8bXAezCKCWhAhguC8l30byv_yWDb-kRRFigxmXKt4XbFL7YlF5WkZc",
		"https://discord.com/api/webhooks/838240872361230387/6Qcxh84DRya7h69Pj8KoNdtBprmpwjWbVkU5ZQ1mOmoJu300_XtvVLxxRwsFrTnw92Qn",
		"https://discord.com/api/webhooks/838240894195990548/3lC7d20pd3pv1lEzwgtpk9eObjt4OEWcikmbk2P0eSlnicENrHefvR5jFzTrk2y5CSiP",
		"https://discord.com/api/webhooks/838240914971426826/snj8osTsgwVSq6Omz2FmbLLQgadgCG_uLaQCzudxzciFarz5ZZZaL-y9n3tDY4nNBpQg",
		"https://discord.com/api/webhooks/838240936562524200/AALqOuo5Zl3iOs1VIp3SWuHnw-ilawYQ3ghjlqnA4q2h3CXAMpmHp91bbbU-Qc8H_TF1",
	}
	webhooks_XP = []string{
		"https://discord.com/api/webhooks/827238648961171517/OPL6TiCE4wuXborIXRpMbP5Jzc-siQvUl-nW1OkU0r1N1qKOc6_jMZhTMeLC0DDSKmZk",
		"https://discord.com/api/webhooks/827238667176509482/Ne0o4bql9PbNEZb7oxXdGKn9K_O_TXjnZJaNvzpR-BXGxxRiNMivnC54G_a8PVNJVU8E",
		"https://discord.com/api/webhooks/827238678563913760/uW4XPy52re8u_sOGRIQXDFumaRi2NukrcQWIAUmenv8kvrs9YSpwyruSBKDJLoyaMax_",
		"https://discord.com/api/webhooks/827238693467062272/G976nfq_3uMesA4XgxkmrUIDdUNw0oTxlhyrd_NU1jgayBagtAJszSgjihuaG9utBgNi",
		"https://discord.com/api/webhooks/827238705378492426/XpkF7zqQxuURgQBkWdFWro7XRMXuRAO5eujjB9v93RZv3Ay4SQkstLxxv-jWZ-JeXBv2",
		"https://discord.com/api/webhooks/827238719403589633/pgcbO0eA6X47d0PdDXR9WSu6o93vSDPsIDqSCyC5viP2kwRnU3-_nlSglAcDDZV-moCy",
		"https://discord.com/api/webhooks/827238735799648276/3eLgdwjxtJYRlPhSXQR-CTVolcYnUTb4UYoVOl1lFhvaCI_qJhIEeOlXG3rFPPqvl6x3",
		"https://discord.com/api/webhooks/827238750542364682/Css0ePFAm0mQ066AtxW413KxMkxpOWAwawn5eLpBt6hNV4nXdXoeJ_nps2u_6s6tmDjp",
		"https://discord.com/api/webhooks/827238764153667594/UW75IAbkt5r6ja2rCZnRtBgdfDUoePCT79kKukKHkASSf0dML14XKgBCwIR7eQ8jFLwp",
		"https://discord.com/api/webhooks/827238779030732830/3doi0zk_7Gsch6aGS6Mik_KUneAYpUEr9Jha6JNyyB6AM8G7b7n4XULg9mwRgtZYCacR",
	}
	webhooks_XA = []string{
		"https://discord.com/api/webhooks/827238276837539851/II6Ahp8_nDueHj5xI0Gf5nn-VoCE2wgNhtet_r8bt1zIznD8MmQdN9ZKgUqGj68R92Pr",
		"https://discord.com/api/webhooks/827238345846947941/H8DjWpYdMjG2KqtTIr6pq9GCUhWbYi6x-xdq24iSKRWwy9EvKP4CX8FShSEHA9jA1gka",
		"https://discord.com/api/webhooks/827238487689527307/FaOFAkZhJW_Xzt86QSqRLu1qr0jI8AOIcSDP0dL8ygk43iFmCqHsbSUpSChByMwqYgFs",
		"https://discord.com/api/webhooks/827238500976820264/Mcm77cq-gk1oxLvH4XHfQUh42gru1KGhnzUM3Izv1R3aiRPzQD4jS4hc0Jps--5jMmC8",
		"https://discord.com/api/webhooks/827238512843030601/1koJatXdWKHTft8hA4IrexNuyS75j5pAG0sGejWX0bvrNaX6YdzKdxdCk4VNv_01muGD",
		"https://discord.com/api/webhooks/827238531516465152/X5sFw3HWpa571VA6xpfr5quXWfCe-lzRqKPEIFlbttAb4wtPow8m60iWcN8NH5ifECGf",
		"https://discord.com/api/webhooks/827238547496239194/T0MlPgm0O7D11kO7w4SlxMFxPtVnf62I88q-IAypxFBqdEz0fOxRi3vJbtQzL1GtnR4n",
		"https://discord.com/api/webhooks/827238563413491733/IifENE3-rTXHxRJbFvUUIXMx8e_ZAsd6lyulcm6an-jWhDHG220wmNvF76WUD_35nOBe",
		"https://discord.com/api/webhooks/827238577170808862/vAgcogKsYvrCxHc6GzAAhyVkUV-O1FwivHhp8pi__y3IrTvHmTZzQjVBGALznSqe3TjI",
		"https://discord.com/api/webhooks/827238592589856828/KLD8VZIUU5No-zreXc66e9T3TKNOs6urK5ti0C8byPG-Ni21cbgiyhUXeaX-cbRb_E01",
	}
	webhooks_EU = []string{
		"https://discord.com/api/webhooks/827239007544016957/-YCobowbVQuPZ8eNkEkwENEzwNDWVj4JFhsmce3-KJ2JKqeqF0misPpepkGzJTSVaSo6",
		"https://discord.com/api/webhooks/827239027262357604/YHsAX1fCtUX8ksWFY6tQ2CXGDY8e4VJwKq3yeFU5EVWdKlvlSmLTYIOpQ66tHkEZhZ_s",
		"https://discord.com/api/webhooks/827239040507838474/CXzTOSTLm1eVbiqhVnD6iHjrSy1eoogERj3RFE1KX6ThducMTxW4zFrWXvTJ2_gQqFYp",
		"https://discord.com/api/webhooks/827239051886460938/OBf-4oGf8bYpMGhBEpMH2zOlYWpgmxp30COI_dFiYgbzRAhPqvF_5xtch2CwY7GTuiPw",
		"https://discord.com/api/webhooks/827239064901517412/Fa4GGaroPqEC7Wn0k5ESEq3aaIwVFACvPyBUFS82eQ_TK0uWtxcn4MynSfcWZSej1LtC",
		"https://discord.com/api/webhooks/827239076226531338/9N-h0puWOPQ9ZBWqsHXsFv4MeTp88ofiGdY_c_ZTWqEjUF5Gdi0kjZzPzWheHVdVTHON",
		"https://discord.com/api/webhooks/827239088573907016/oGFf8-3ecNkg78vqNCR6a61KzdxXwfNMFmWL5VN5Nn6Z-PiyrXR-OjL0VzuC092ILNcg",
		"https://discord.com/api/webhooks/827239101513990214/TQO8rEcdLGIGkab-C2VH3qwSFKYMZRlZrcUhEv-0-A3eDJf1IdulR8hd4SvDhkzQJvmY",
		"https://discord.com/api/webhooks/827239114620796928/8w1aA8HSGRl0GDowO7Z60WNAaIUbpmtUV0Zv5FIqeOI_scDy_HhmbWtQuKgHDFpoUPdA",
		"https://discord.com/api/webhooks/827239127929716857/wXR2jMQxu2YldlM221QWahU9zMMoHCUulxlmlazBEJkQ1ofy_tRmfIPCIWpIlKfgrEup",
	}
	webhooks_US = []string{
		"https://discord.com/api/webhooks/827238832919019580/UoKzuExTvKblGPiYI4LCxlqpsKG55TAeNClXmMQJxx9spqiX35fwei8e0P1YnA3A8dG0",
		"https://discord.com/api/webhooks/827238853332303883/cDOvVWER47ZMTWAywUfMPLVpbuQcLjhzH-SVYpAKVa8hHr1Hh_Qgn60NeoG5LyoG-2Se",
		"https://discord.com/api/webhooks/827238867739738142/SgBcbs2CAznPpo0WlSM_3zk_e711lj26s6kZnBcXC_03ovFKtE8BrKeylTDIbmg8TKbv",
		"https://discord.com/api/webhooks/827238881224556577/nJ5WlNZnL7K4mNAuRzaPhVkpy70iSr3Yq_uWlnYx_cvHppI8a6KH8b1w4ZVYVJRckpvG",
		"https://discord.com/api/webhooks/827238894939799614/XofVlqGaF3-tLev_hjkx9NfS9yLn7-Tmso-cdo32nAtGOSAUvFxw2ia4XR5vp9ouzWyR",
		"https://discord.com/api/webhooks/827238907430568015/Q6GDyFf06iOGxsFTF9paobMuHFH7qd0ai8m3l7ah3uVV3Rlj_jADuYTSEJKGd4ZLPLNu",
		"https://discord.com/api/webhooks/827238919632453653/20z-A-XUB47iCVaSmIx8Se18FvH5Kuropa6MjrceecPvszVPBL9VXJDYpDtd5nTu1-11",
		"https://discord.com/api/webhooks/827238933489385493/uyjZU0imOLAxRUxVfNj6CzFpQ87It4cWC5GZ1KKKLdmgND4-oQDBkF5A8LY-AlyZNGQu",
		"https://discord.com/api/webhooks/827238947817127946/VZz2HgXRR59N480nEoMytPPZ5OF18PiOwvE6iveNWaVVJv2PGjALRc3nITOl4Fvva8Za",
		"https://discord.com/api/webhooks/827238962690129940/dboP3T6dFkBYP_dUAqPSK4ukW5-RqMj-xeBYE6fkki53AJctWjpk5gBR7gDDu0hxDYi1",
	}
	webhooks_CN = []string{
		"https://discord.com/api/webhooks/827239173970591744/c6xyGZduDBYAFvy1vm-iSAhVRQlZDoDOPY_KQC9IOQ-0nb1VnihH4ZpML15Wg-jLX4I-",
		"https://discord.com/api/webhooks/827239192526192710/Kh4wHlVuY_djj_W9yx77XjABKudPLdZHk6wPi2EIDXsSnulm_JufNHd0djnZnWRFb5x4",
		"https://discord.com/api/webhooks/827239205671010355/2YfE7NQFcfKXIx2YiDcYjlaNtiVLki4Ic_RMy3qupCxYCSa2URYhwcEBgyFVxG04m6Kx",
		"https://discord.com/api/webhooks/827239216671883284/FiMN-U7y73rLKb8m0Vbl6PVu2ryULKJV91LgOCoqc1Dyd1Pbe_pfXTttKC5XCls0csH4",
		"https://discord.com/api/webhooks/827239229305389067/OTwwoqJFMkZAQLfaYPOKsPX9GKU-mZR8wtkVYS64hG3w6fOUay5W5hpashSdfRkls7ms",
		"https://discord.com/api/webhooks/827239242659921920/povK2ryNhk8oeIy9UPCrhqBSpLS17tYXdqxY5329yAsCREhU_FHdBpBksdH__Y6Epwxy",
		"https://discord.com/api/webhooks/827595497216147549/APD9eGl6cOq-1rXlK1iLQR8jR7pM0G4abZnWMHB8HQ_j4kjRQKmjHVi2A4Hys7aYPj6L",
		"https://discord.com/api/webhooks/827595562928177172/4zjjqiji_APfa-yA4Xr6PFweHcGeQgWyq2tLADLRH8PSPLI0dOPTPBT0xRxShR7frTXw",
		"https://discord.com/api/webhooks/827595577561841665/9aqCKm0NQXvmYjIbzi5sWvVr5E2Je8zKqRKmRnU1hD3bsMxlKUq3dw7PvRLcZzyKnO6y",
		"https://discord.com/api/webhooks/827595593035022346/TZPrK_wkRyIMZT6nrFgoeB9Qtm49qJ36hgjz0PMBGNZDTrdOo8hffMAaNtYOSJJdFoJV",
	}
	webhooks_JP = []string{
		"https://discord.com/api/webhooks/827595633367449651/VAGYtK62p5DNHZprwb9qUB5NxMjziPC5bsNNJOZeqirbvjV0m_PtXyjaZWE6OMBh93gx",
		"https://discord.com/api/webhooks/827595653302452224/ZGnxOLoHbfKnS7D82IvP1FNv7DLO2X-eXETqgCPa1KhbI64-7Gk6KPtZOnU0k4xto90s",
		"https://discord.com/api/webhooks/827595666564972605/j4bXidwGi4TWmPLjamNBZitxpKUyH0nyIzO08_JzFEXCjwJWEv1NKrOzNZ6H2nsIRi-m",
		"https://discord.com/api/webhooks/827595678086987857/pLLU0Q8Fn_3xqwfjnUPp_2h8ci-0T-lk6IiuaHbl_8ve8NPYdv0KYzQi_ylEQkdsDsrf",
		"https://discord.com/api/webhooks/827595691840766002/DcGl_gkA_BgezAcbfBbWudBfTg-Gsb8pOJp5TKm_1mFhTzPB5Gd0soLG-TOitgKuhbYy",
		"https://discord.com/api/webhooks/827595703223058512/FpOWThOqruHVMF6RwnRDKD3h8pXW5r_bxcQnqqC4I1_WhX7I0SYsgCwOwerEtuBiDs9g",
		"https://discord.com/api/webhooks/827595716817977364/CdB9Xwby6SVf0EvOJ91g7NNGJn7zro3w3_ygEL_Vd1ImiiTxKUugL7mkJLgPkZhxAVkb",
		"https://discord.com/api/webhooks/827595730058870804/NkRA8Y9QPorWjwShun38xHn2N0UEV1NOuwfTqrJrmdh1Fa_sAapKFrIoxM-m1GaECPdk",
		"https://discord.com/api/webhooks/827595744562249728/vY77xeO1tFJPonPsCBDk6oyf8GyjR7T63gKBkvrkm1-h6AtyM0fqpW181EY0pEi2zybp",
		"https://discord.com/api/webhooks/827595759456092230/fa9evP4eHNt5iLyUxCbGP2GlVMtZbHFvVdXQbFM_rHMmYRDPeiSYgwd1br_ZMA5EZTFU",
	}
	webhooks_SSENSE = []string{
		"https://discord.com/api/webhooks/853644919034216458/_s-hxFYm1HO5yez27MtNpB8WNvc01I9QcPd22En-18Om8HHy4KlQqAWnyLpTrSqd9MZ8",
		"https://discord.com/api/webhooks/853645171904086076/27niQUR1YawaVKJ-LZNuWJfcO_IS83ucPKDpeJpm698jLWwEyxXLa-nlnsaAnk7R3rtU",
		"https://discord.com/api/webhooks/853645202082103327/VEKiofNMo-4oMlub3HKnwTuM2Ky4NK84_7KHFTq2jUVnzemZ7jzsZ2MSVR1kbVmN-zry",
		"https://discord.com/api/webhooks/853645221119393832/a8jxXRW9AUIu3Ng6ichSuHJ5GHd6XM2sbhyT589IlKkBQJQKU1hvwu4yEer5YYIk1R-U",
		"https://discord.com/api/webhooks/853645240711381027/T3jFb3r-cYD0WTP74lqRO04DIh2UhQjT7J3TBdRw0JkvXERZZgAZwSqmDwcOHei0hazE",
		"https://discord.com/api/webhooks/853645265914036244/iTFmV-At-MqsqltOyYjO_QM0TQGQzK976kl2g02uvEAzWFNOkdqc0Dmv1uH48wX4LFDO",
		"https://discord.com/api/webhooks/853645299557466122/S9ApYktYYEOU0W25zMTIbDUw-Jik_6QRe5D1uX8ez-2UbkVHtzFEEFApvQql3M0-F1g0",
		"https://discord.com/api/webhooks/853645320416002078/8YvFCuyXE5D7JDybHx6o1WXWrFLBKhpDZuUBq-yAmWyPP4VrX7q6DogYWaxa_OpE-IxN",
		"https://discord.com/api/webhooks/853645340145483776/tnaCzWG62987QRtEEI0IulVtHLXUVWpu2kGF9w_7ZklrBgsf8QafjXsx0z1Lhrj2UIhR",
		"https://discord.com/api/webhooks/853645361989156864/J90hdO3RRtaw30rO49AuHtcMh_KcpK_2yikS1ZODKr5xc8UXPq7RVj40u5y_BbefmPV0",
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
