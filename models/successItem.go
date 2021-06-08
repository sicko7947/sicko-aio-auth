package models

import (
	auth_service "github.com/sicko7947/sicko-aio-auth/proto/auth"
)

type SuccessItem struct {
	KeyId   string
	Setup   *auth_service.SuccessSetup
	Product *auth_service.SuccessProduct
}
