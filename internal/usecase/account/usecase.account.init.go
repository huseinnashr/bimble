package account

import "github.com/huseinnashr/bimble/internal/domain"

type Usecase struct {
}

func New(accountRepo domain.IAccountRepo) *Usecase {
	return &Usecase{}
}
