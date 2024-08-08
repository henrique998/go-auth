package repositories

import (
	"database/sql"

	"github.com/henrique998/go-auth/internal/app/entities"
)

type PGLoginAttemptsRepository struct {
	Db *sql.DB
}

func (r *PGLoginAttemptsRepository) Create(la entities.LoginAttempt) error {
	query :=
		`INSERT INTO login_attempts (id, account_email, ip, user_agent, success, attempted_at)
		VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := r.Db.Exec(query,
		la.GetID(),
		la.GetEmail(),
		la.GetIP(),
		la.GetUserAgent(),
		la.GetSuccess(),
		la.GetAttemptedAt(),
	)
	if err != nil {
		return err
	}

	return nil
}
