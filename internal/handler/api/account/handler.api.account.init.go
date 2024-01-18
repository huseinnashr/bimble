package account

import (
	v1 "github.com/huseinnashr/bimble/api/v1"
	"github.com/huseinnashr/bimble/internal/domain"
)

type Handler struct {
	v1.UnimplementedAccountServiceServer
	accountUsecase domain.IAccountUsecase
}

func New(accountUsecase domain.IAccountUsecase) v1.AccountServiceServer {
	return &Handler{
		accountUsecase: accountUsecase,
	}
}
