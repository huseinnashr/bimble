package account

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/huseinnashr/bimble/internal/domain"
)

func (r *Repo) CreateAccount(ctx context.Context, email, hashedPassword string) (int64, error) {
	row := r.sqlDatabase.QueryRowContext(ctx,
		`
			INSERT INTO accounts(email, password) VALUES ($1,$2) 
			ON CONFLICT ON CONSTRAINT accounts_email_key
			DO UPDATE SET password = EXCLUDED.password
			RETURNING id
		`,
		email, hashedPassword,
	)
	if err := row.Err(); err != nil {
		return 0, err
	}

	var accountID int64
	if err := row.Scan(&accountID); err != nil {
		return 0, err
	}

	return accountID, nil
}

func (r *Repo) SetVerificationToken(ctx context.Context, accountID int64, token string) error {
	response := r.redis.Set(ctx, token, accountID, 1*time.Hour)
	if err := response.Err(); err != nil {
		return err
	}

	return nil
}

func (r *Repo) SendVerificationToken(ctx context.Context, email, encodedToken string) error {
	verificationLink := fmt.Sprintf("http://localhost:%d/accounts:verify?token=%s", r.config.Port, encodedToken)
	log.Println(verificationLink, "send to", email)

	return nil
}

func (r *Repo) GetAccountIDFromToken(ctx context.Context, token string) (int64, error) {
	response := r.redis.Get(ctx, token)
	if err := response.Err(); err != nil {
		return 0, err
	}

	accountID, err := response.Int64()
	if err != nil {
		return 0, err
	}

	return accountID, nil
}

func (r *Repo) SetAccountToVerified(ctx context.Context, accountID int64) error {
	_, err := r.sqlDatabase.ExecContext(ctx, `UPDATE accounts SET is_verified = true WHERE id = $1`, accountID)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) GetAccountRefFromEmail(ctx context.Context, email string) (*domain.AccountRef, error) {
	row := r.sqlDatabase.QueryRowContext(ctx,
		`
			SELECT 
				accounts.id, COALESCE(profiles.id, 0), COALESCE(preferences.id, 0), 
				accounts.password, accounts.is_verified
			FROM accounts 
			LEFT JOIN profiles ON profiles.account_id = accounts.id AND profiles.type = 'DATING'
			LEFT JOIN preferences ON preferences.profile_id = profiles.id
			WHERE accounts.email = $1
			LIMIT 1
		`,
		email,
	)
	if err := row.Err(); err != nil {
		return nil, err
	}

	accountRef := &domain.AccountRef{}
	err := row.Scan(
		&accountRef.AccountID, &accountRef.DatingProfileID, &accountRef.DatingPreferenceID,
		&accountRef.HashedPassword, &accountRef.IsVerified,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return accountRef, nil
}

func (r *Repo) SetSession(ctx context.Context, token string, accountRef domain.AccountRef) error {
	key := fmt.Sprintf("login_session:%d", accountRef.AccountID)

	val, _ := json.Marshal(accountRef)

	response := r.redis.Set(ctx, key, val, 24*time.Hour)
	if err := response.Err(); err != nil {
		return err
	}

	return nil
}
