package account

import (
	v1 "github.com/huseinnashr/bimble/api/v1"
	"github.com/huseinnashr/bimble/internal/domain"
)

type Handler struct {
	accountUsecase domain.IAccountUsecase
}

func New(accountUsecase domain.IAccountUsecase) v1.AccountServiceHTTPServer {
	return &Handler{
		accountUsecase: accountUsecase,
	}
}
