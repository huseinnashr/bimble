package account

import (
	"github.com/huseinnashr/bimble/internal/domain"
)

type Repo struct {
}

func New() domain.IAccountRepo {
	return &Repo{}
}
