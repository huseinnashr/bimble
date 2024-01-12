package account

import "github.com/huseinnashr/bimble/internal/domain"

type Usecase struct {
	accountRepo domain.IAccountRepo
}

func New(accountRepo domain.IAccountRepo) domain.IAccountUsecase {
	return &Usecase{
		accountRepo: accountRepo,
	}
}
