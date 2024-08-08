package repositories

import (
	"database/sql"
	"time"

	"github.com/henrique998/go-auth/internal/app/entities"
	"github.com/henrique998/go-auth/internal/configs/logger"
)

type PGMagicLinksRepository struct {
	Db *sql.DB
}

type magicLinkRecord struct {
	id        string
	accountId string
	code      string
	expiresAt time.Time
	createdAt time.Time
}

func (r *PGMagicLinksRepository) FindByValue(val string) entities.MagicLink {
	var magicLinkData magicLinkRecord

	query := "SELECT id, account_id, code, expires_at, created_at FROM magic_links WHERE code = $1"
	row := r.Db.QueryRow(query, val)

	err := row.Scan(&magicLinkData.id, &magicLinkData.accountId, &magicLinkData.code, &magicLinkData.expiresAt, &magicLinkData.createdAt)
	if err != nil {
		if err != sql.ErrNoRows {
			logger.Error("Error trying to find verification code", err)
		}
		return nil
	}

	magicLink := entities.NewExistingMagicLink(
		magicLinkData.id,
		magicLinkData.code,
		magicLinkData.accountId,
		magicLinkData.expiresAt,
		magicLinkData.createdAt,
	)

	return magicLink
}

func (r *PGMagicLinksRepository) Create(ml entities.MagicLink) error {
	query := "INSERT INTO magic_links (id, account_id, code, expires_at, created_at) VALUES ($1, $2, $3, $4, $5)"

	_, err := r.Db.Exec(query,
		ml.GetID(),
		ml.GetAccountID(),
		ml.GetCode(),
		ml.GetExpiresAt(),
		ml.GetCreatedAt(),
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *PGMagicLinksRepository) Delete(mId string) error {
	query := "DELETE FROM magic_links WHERE id = $1"

	_, err := r.Db.Exec(query, mId)
	if err != nil {
		return err
	}

	return nil
}
