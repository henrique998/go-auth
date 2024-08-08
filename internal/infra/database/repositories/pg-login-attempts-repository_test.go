package repositories

import (
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/henrique998/go-auth/internal/app/entities"
	"github.com/stretchr/testify/assert"
)

func TestPGLoginAttemptsRepository_Create(t *testing.T) {
	assert := assert.New(t)

	db, mock, err := sqlmock.New()
	assert.NoError(err)
	defer db.Close()

	repo := PGLoginAttemptsRepository{Db: db}

	data := entities.NewExistingLoginAttempt(
		"fake-attempt-id",
		"jhondoe@email.com",
		"fake-attempt-ip",
		"fake-user-agent",
		true,
		time.Now(),
	)

	t.Run("Success", func(t *testing.T) {
		mock.ExpectExec(`INSERT INTO login_attempts \(id, account_email, ip, user_agent, success, attempted_at\)
		VALUES \(\$1, \$2, \$3, \$4, \$5, \$6\)`).
			WithArgs(
				data.GetID(),
				data.GetEmail(),
				data.GetIP(),
				data.GetUserAgent(),
				data.GetSuccess(),
				data.GetAttemptedAt(),
			).WillReturnResult(sqlmock.NewResult(1, 1))

		err = repo.Create(data)

		assert.NoError(err)

		err = mock.ExpectationsWereMet()
		assert.NoError(err)
	})

	t.Run("Error", func(t *testing.T) {
		mock.ExpectExec(`INSERT INTO login_attempts \(id, account_email, ip, user_agent, success, attempted_at\)
		VALUES \(\$1, \$2, \$3, \$4, \$5, \$6\)`).
			WithArgs(
				data.GetID(),
				data.GetEmail(),
				data.GetIP(),
				data.GetUserAgent(),
				data.GetSuccess(),
				data.GetAttemptedAt(),
			).WillReturnError(errors.New("insert failed"))

		err = repo.Create(data)

		assert.Error(err)
		assert.Equal("insert failed", err.Error())

		err = mock.ExpectationsWereMet()
		assert.NoError(err)
	})
}
