package account

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/google/uuid"
)

func (u *Usecase) Signup(ctx context.Context, email, password string) error {
	hashedPassword, err := u.accountRepo.HashPassword(ctx, password)
	if err != nil {
		return errors.InternalServer(err.Error(), "failed to hash password")
	}

	accountID, err := u.accountRepo.CreateAccount(ctx, email, hashedPassword)
	if err != nil {
		return errors.InternalServer(err.Error(), "failed to create account")
	}

	uuid, err := uuid.NewRandom()
	if err != nil {
		return errors.InternalServer(err.Error(), "failed to create verification token")
	}

	token := fmt.Sprintf("%d:%s", accountID, uuid.String())
	if err := u.accountRepo.SetVerificationToken(ctx, accountID, token); err != nil {
		return errors.InternalServer(err.Error(), "failed to set verification token")
	}

	encodedToken := base64.URLEncoding.EncodeToString([]byte(token))
	if err := u.accountRepo.SendVerificationToken(ctx, email, encodedToken); err != nil {
		return errors.InternalServer(err.Error(), "failed to send verification token")
	}

	return nil
}

func (u *Usecase) Verify(ctx context.Context, encodedToken string) error {
	token, err := base64.URLEncoding.DecodeString(encodedToken)
	if err != nil {
		return errors.InternalServer(err.Error(), "failed to decode token")
	}

	accountID, err := u.accountRepo.GetAccountIDFromToken(ctx, string(token))
	if err != nil {
		return errors.InternalServer(err.Error(), "failed to get account id")
	}

	if accountID == 0 {
		return errors.BadRequest("account id is not found", "Verification link is invalid")
	}

	if err := u.accountRepo.SetAccountToVerified(ctx, accountID); err != nil {
		return errors.InternalServer(err.Error(), "failed to set account to verified")
	}

	return nil
}

func (u *Usecase) Login(ctx context.Context, email, password string) (string, error) {
	accountRef, err := u.accountRepo.GetAccountRefFromEmail(ctx, email)
	if err != nil {
		return "", errors.InternalServer(err.Error(), "failed to retrieve account ref")
	}

	hashedPassword, err := u.accountRepo.HashPassword(ctx, password)
	if err != nil {
		return "", errors.InternalServer(err.Error(), "failed to hash password")
	}

	if hashedPassword != accountRef.HashedPassword {
		return "", errors.Unauthorized("hashed password not matched", "email/password is invalid")
	}

	uuid, err := uuid.NewRandom()
	if err != nil {
		return "", errors.InternalServer(err.Error(), "failed to create verification token")
	}

	token := fmt.Sprintf("%d:%s", accountRef.AccountID, uuid.String())
	if err := u.accountRepo.SetSession(ctx, token, accountRef); err != nil {
		return "", errors.InternalServer(err.Error(), "failed to set session")
	}

	encodedToken := base64.URLEncoding.EncodeToString([]byte(token))
	return encodedToken, nil
}
