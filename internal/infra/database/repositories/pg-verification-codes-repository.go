package repositories

import (
	"database/sql"
	"time"

	"github.com/henrique998/go-auth/internal/app/entities"
	"github.com/henrique998/go-auth/internal/configs/logger"
)

type PGVerificationCodesRepository struct {
	Db *sql.DB
}

type VerificationCodeRecord struct {
	ID        string
	AccountId string
	Value     string
	ExpiresAt time.Time
	CreatedAt time.Time
}

func (r *PGVerificationCodesRepository) FindByValue(val string) entities.VerificationCode {
	var verificationCodeData VerificationCodeRecord

	query := "SELECT id, account_id, value, expires_at, created_at FROM verification_codes WHERE value = $1"
	row := r.Db.QueryRow(query, val)

	err := row.Scan(&verificationCodeData.ID, &verificationCodeData.AccountId, &verificationCodeData.Value, &verificationCodeData.CreatedAt, &verificationCodeData.ExpiresAt)
	if err != nil {
		if err != sql.ErrNoRows {
			logger.Error("Error trying to find verification code", err)
		}
		return nil
	}

	verificationCode := entities.NewExistingVerificationCode(
		verificationCodeData.ID,
		verificationCodeData.Value,
		verificationCodeData.AccountId,
		verificationCodeData.ExpiresAt,
		verificationCodeData.CreatedAt,
	)

	return verificationCode
}

func (r *PGVerificationCodesRepository) Create(vc entities.VerificationCode) error {
	query := "INSERT INTO verification_codes (id, account_id, value, expires_at, created_at) VALUES($1, $2, $3, $4, $5)"

	_, err := r.Db.Exec(query,
		vc.GetID(),
		vc.GetAccountID(),
		vc.GetValue(),
		vc.GetExpiresAt(),
		vc.GetCreatedAt(),
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *PGVerificationCodesRepository) Delete(codeId string) error {
	query := "DELETE FROM verification_codes WHERE id = $1"

	_, err := r.Db.Exec(query, codeId)
	if err != nil {
		return err
	}

	return nil
}
