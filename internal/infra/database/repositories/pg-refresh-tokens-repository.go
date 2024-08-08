package repositories

import (
	"database/sql"
	"time"

	"github.com/henrique998/go-auth/internal/app/entities"
	"github.com/henrique998/go-auth/internal/configs/logger"
)

type PGRefreshTokensRepository struct {
	Db *sql.DB
}

type refreshTokenRecord struct {
	id        string
	value     string
	accountId string
	expiresAt time.Time
	createdAt time.Time
}

func (r *PGRefreshTokensRepository) FindByValue(val string) entities.RefreshToken {
	var refreshTokenData refreshTokenRecord

	query := "SELECT id, value, account_id, expires_at, created_at FROM refresh_tokens WHERE refresh_token = $1"
	row := r.Db.QueryRow(query, val)

	err := row.Scan(&refreshTokenData.id, &refreshTokenData.value, &refreshTokenData.accountId, &refreshTokenData.expiresAt, &refreshTokenData.createdAt)
	if err != nil {
		if err != sql.ErrNoRows {
			logger.Error("Error trying to find refresh token", err)
		}
		return nil
	}

	refreshToken := entities.NewExistingRefreshToken(
		refreshTokenData.id,
		refreshTokenData.value,
		refreshTokenData.accountId,
		refreshTokenData.expiresAt,
		refreshTokenData.createdAt,
	)

	return refreshToken
}

func (r *PGRefreshTokensRepository) Create(rt entities.RefreshToken) error {
	query := "INSERT INTO refresh_tokens (id, value, account_id, expires_at, created_at) VALUES($1, $2, $3, $4, $5)"

	_, err := r.Db.Exec(query,
		rt.GetID(),
		rt.GetValue(),
		rt.GetAccountID(),
		rt.GetExpiresAt(),
		rt.GetCreatedAt(),
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *PGRefreshTokensRepository) Delete(val string) error {
	query := "DELETE FROM refresh_tokens WHERE value = $1"

	_, err := r.Db.Exec(query, val)
	if err != nil {
		return err
	}

	return nil
}
