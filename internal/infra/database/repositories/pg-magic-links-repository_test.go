package repositories

import (
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/henrique998/go-auth/internal/app/entities"
	"github.com/henrique998/go-auth/internal/infra/utils"
	"github.com/stretchr/testify/assert"
)

func TestPGMagicLinksRepository_FindByValue(t *testing.T) {
	assert := assert.New(t)

	db, mock, err := sqlmock.New()
	assert.NoError(err)
	defer db.Close()

	repo := PGMagicLinksRepository{Db: db}

	id := "fake-magic-link-id"
	code, _ := utils.GenerateCode(10)
	expiresAt := time.Now().Add(15 * time.Minute)

	data := entities.NewExistingMagicLink(id, "fake-account-id", code, expiresAt, time.Now())

	t.Run("Success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{
			"id", "account_id", "code", "expires_at", "created_at",
		}).AddRow(
			data.GetID(),
			data.GetAccountID(),
			data.GetCode(),
			data.GetExpiresAt(),
			data.GetCreatedAt(),
		)

		mock.ExpectQuery(`SELECT id, account_id, code, expires_at, created_at FROM magic_links WHERE code = \$1`).
			WithArgs(data.GetCode()).WillReturnRows(rows)

		link := repo.FindByValue(data.GetCode())

		assert.NotNil(link)
		assert.Equal(link.GetID(), id)
		assert.Equal(link.GetExpiresAt(), data.GetExpiresAt())

		err = mock.ExpectationsWereMet()
		assert.NoError(err)
	})

	t.Run("Error", func(t *testing.T) {
		mock.ExpectQuery(`SELECT id, account_id, code, expires_at, created_at FROM magic_links WHERE code = \$1`).
			WithArgs(data.GetCode()).WillReturnError(errors.New("some error"))

		link := repo.FindByValue(data.GetCode())

		assert.Nil(link)

		err = mock.ExpectationsWereMet()
		assert.NoError(err)
	})
}

func TestPGMagicLinksRepository_Create(t *testing.T) {
	assert := assert.New(t)

	db, mock, err := sqlmock.New()
	assert.NoError(err)
	defer db.Close()

	repo := PGMagicLinksRepository{Db: db}

	code, _ := utils.GenerateCode(10)

	data := entities.NewExistingMagicLink(
		"fake-magic-link-id",
		"fake-account-id",
		code,
		time.Now().Add(15*time.Minute),
		time.Now(),
	)

	t.Run("Success", func(t *testing.T) {
		mock.ExpectExec(`INSERT INTO magic_links \(id, account_id, code, expires_at, created_at\) VALUES \(\$1, \$2, \$3, \$4, \$5\)`).
			WithArgs(
				data.GetID(),
				data.GetAccountID(),
				data.GetCode(),
				data.GetExpiresAt(),
				data.GetCreatedAt(),
			).WillReturnResult(sqlmock.NewResult(1, 1))

		err := repo.Create(data)

		assert.NoError(err)

		err = mock.ExpectationsWereMet()
		assert.NoError(err)
	})

	t.Run("Error", func(t *testing.T) {
		mock.ExpectExec(`INSERT INTO magic_links \(id, account_id, code, expires_at, created_at\) VALUES \(\$1, \$2, \$3, \$4, \$5\)`).
			WithArgs(
				data.GetID(),
				data.GetAccountID(),
				data.GetCode(),
				data.GetExpiresAt(),
				data.GetCreatedAt(),
			).WillReturnError(errors.New("insert failed"))

		err := repo.Create(data)

		assert.Error(err)
		assert.Equal("insert failed", err.Error())

		err = mock.ExpectationsWereMet()
		assert.NoError(err)
	})
}

func TestPGMagicLinksRepository_Delete(t *testing.T) {
	assert := assert.New(t)

	db, mock, err := sqlmock.New()
	assert.NoError(err)
	defer db.Close()

	repo := PGMagicLinksRepository{Db: db}

	id := "fake-magic-link-id"

	t.Run("Success", func(t *testing.T) {
		mock.ExpectExec(`DELETE FROM magic_links WHERE id = \$1`).WithArgs(id).WillReturnResult(sqlmock.NewResult(1, 1))

		err := repo.Delete(id)

		assert.NoError(err)

		err = mock.ExpectationsWereMet()
		assert.NoError(err)
	})

	t.Run("Error", func(t *testing.T) {
		mock.ExpectExec(`DELETE FROM magic_links WHERE id = \$1`).
			WithArgs(id).
			WillReturnError(errors.New("delete failed"))

		err := repo.Delete(id)

		assert.Error(err)
		assert.Equal("delete failed", err.Error())

		err = mock.ExpectationsWereMet()
		assert.NoError(err)
	})
}
