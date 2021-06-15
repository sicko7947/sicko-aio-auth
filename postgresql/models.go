package postgresql

import (
	"time"
)

type STATUSCODE int64

const (
	OK                     STATUSCODE = 200
	NOT_FOUND              STATUSCODE = 404
	DATABASE_ERROR         STATUSCODE = 400
	CANCELLED              STATUSCODE = 1001
	ACTIVATED              STATUSCODE = 1002
	KEY_ERROR              STATUSCODE = 1003
	LOGGED_IN_OTHER_DEVICE STATUSCODE = 1005
	NO_ACCESS              STATUSCODE = 1007
	KEY_STATUS_NOT_MATCH   STATUSCODE = 1008
	WRONG_CREDENTIALS      STATUSCODE = 1009
	REQUIRE_ACTIVATION     STATUSCODE = 1010
)

var (
	STATUSMAP map[STATUSCODE]string = map[STATUSCODE]string{
		OK:                     "Success.",
		NOT_FOUND:              "Key Not Found.",
		DATABASE_ERROR:         "Internal Server Error",
		CANCELLED:              "The key has Suspended.",
		ACTIVATED:              "This key is activated.",
		KEY_ERROR:              "Key Error.",
		LOGGED_IN_OTHER_DEVICE: "The key is already logged in on other devices.",
		NO_ACCESS:              "The key has no access.",
		KEY_STATUS_NOT_MATCH:   "The key status doesn't match the record.",
		WRONG_CREDENTIALS:      "Wrong KEY or Email value",
		REQUIRE_ACTIVATION:     "Activate Your Key before Login.",
	}
)

type keyMain struct {
	Id           int64
	Key          string
	Status       int64
	Email        string
	DiscordID    string
	CreateTime   time.Time
	ActivateTime time.Time
	ExpireTime   time.Time
	KeyType      string
	Reason       string
}

type keyDetails struct {
	Id            int64
	Key           string
	IP            string
	CpuId         string
	LastLoginTime time.Time
}

type successTable struct {
	Id              int64
	KeyId           string
	EntryId         string
	Timestamp       time.Time
	TaskType        string
	MonitorMode     string
	ProductId       string
	Category        string
	Region          string
	UsePsychoCookie bool
	OtherC          string
	OtherD          string
	OtherE          string
}

type productDetail struct {
	Id                 int64
	ProductId          string
	MerchGroup         string
	ProductSku         string
	ProductName        string
	ProductDescription string
	Size               string
	Price              string
	Quantity           int64
	OrderNumber        string
	ProfileName        string
	Email              string
	Account            string
	GiftCards          string
	DiscountCode       string
	ImageUrl           string
	RedirectUrl        string
	OtherA             string
	OtherB             string
	OtherC             string
}

type successItem struct {
	Category    string
	Region      string
	ProductSku  string
	ProductName string
	OrderNumber string
	Email       string
	Size        string
	Timestamp   time.Time
	ImageUrl    string
	RedirectUrl string
}
