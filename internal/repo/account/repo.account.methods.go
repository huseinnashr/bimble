package account

import (
	"context"

	"github.com/huseinnashr/bimble/internal/domain"
)

func (r *Repo) HashPassword(ctx context.Context, password string) (string, error) {
	return "", nil
}

func (r *Repo) CreateAccount(ctx context.Context, email, hashedPassword string) (int64, error) {
	return 0, nil
}

func (*Repo) SetVerificationToken(ctx context.Context, accountID int64, encodedToken string) error {
	return nil
}

func (r *Repo) SendVerificationToken(ctx context.Context, email, token string) error {
	return nil
}

func (r *Repo) GetAccountIDFromToken(ctx context.Context, token string) (int64, error) {
	return 0, nil
}

func (r *Repo) SetAccountToVerified(ctx context.Context, accountID int64) error {
	return nil
}

func (r *Repo) GetAccountRefFromEmail(ctx context.Context, email string) (domain.AccountRef, error) {
	return domain.AccountRef{}, nil
}

func (*Repo) SetSession(ctx context.Context, token string, accountRef domain.AccountRef) error {
	return nil
}
