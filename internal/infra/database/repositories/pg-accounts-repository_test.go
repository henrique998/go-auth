package repositories

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/henrique998/go-auth/internal/app/entities"
	"github.com/stretchr/testify/assert"
)

func TestPGAccountsRepository_FindById(t *testing.T) {
	assert := assert.New(t)

	db, mock, err := sqlmock.New()
	assert.NoError(err)
	defer db.Close()

	name := "jhon"
	email := "jhondoe@gmail.com"
	hashedPass := "hashed-pass"
	phone := "999999999"
	var age int8 = 23

	accountData := entities.NewAccount(name, email, hashedPass, phone, age, "")

	t.Run("Success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{
			"id", "name", "email", "password_hash", "phone", "age", "provider_id", "is_2fa_enabled", "is_email_verified", "last_login_at", "last_login_ip", "last_login_country", "last_login_city", "created_at", "updated_at",
		}).AddRow(
			accountData.GetID(),
			accountData.GetName(),
			accountData.GetEmail(),
			accountData.GetPass(),
			accountData.GetPhone(),
			accountData.GetAge(),
			accountData.GetProviderId(),
			accountData.GetIs2FaEnabled(),
			accountData.GetIsEmailVerified(),
			accountData.GetLastLoginAt(),
			accountData.GetLastLoginIp(),
			accountData.GetLastLoginCountry(),
			accountData.GetLastLoginCity(),
			accountData.GetCreatedAt(),
			accountData.GetUpdatedAt(),
		)

		mock.ExpectQuery("SELECT id, name, email, password_hash, phone, age, provider_id, is_2fa_enabled, is_email_verified, last_login_at, last_login_ip, last_login_country, last_login_city, created_at, updated_at FROM accounts WHERE id = \\$1 LIMIT 1").WithArgs(accountData.GetID()).WillReturnRows(rows)

		repo := PGAccountsRepository{
			Db: db,
		}

		account := repo.FindById(accountData.GetID())

		assert.NotNil(account)
		assert.Equal(accountData.GetID(), account.GetID())
		assert.Equal(accountData.GetEmail(), account.GetEmail())

		err = mock.ExpectationsWereMet()
		assert.NoError(err)
	})

	t.Run("ErrorNotNoRows", func(t *testing.T) {
		mock.ExpectQuery("SELECT id, name, email, password_hash, phone, age, provider_id, is_2fa_enabled, is_email_verified, last_login_at, last_login_ip, last_login_country, last_login_city, created_at, updated_at FROM accounts WHERE id = \\$1 LIMIT 1").
			WithArgs("fake-id").
			WillReturnError(errors.New("some other error"))

		repo := PGAccountsRepository{
			Db: db,
		}

		account := repo.FindById("fake-id")

		assert.Nil(account)

		err = mock.ExpectationsWereMet()
		assert.NoError(err)
	})

}

func TestPGAccountsRepository_FindByEmail(t *testing.T) {
	assert := assert.New(t)

	db, mock, err := sqlmock.New()
	assert.NoError(err)
	defer db.Close()

	email := "jhondoe@email.com"
	hashedPass := "hashed-pass"
	phone := "999999999"
	providerId := ""

	accountData := entities.NewAccount("jhondoe", email, hashedPass, phone, 23, providerId)

	t.Run("Success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{
			"id", "name", "email", "password_hash", "phone", "age", "provider_id", "is_2fa_enabled", "is_email_verified", "last_login_at", "last_login_ip", "last_login_country", "last_login_city", "created_at", "updated_at",
		}).AddRow(
			accountData.GetID(),
			accountData.GetName(),
			accountData.GetEmail(),
			accountData.GetPass(),
			accountData.GetPhone(),
			accountData.GetAge(),
			accountData.GetProviderId(),
			accountData.GetIs2FaEnabled(),
			accountData.GetIsEmailVerified(),
			accountData.GetLastLoginAt(),
			accountData.GetLastLoginIp(),
			accountData.GetLastLoginCountry(),
			accountData.GetLastLoginCity(),
			accountData.GetCreatedAt(),
			accountData.GetUpdatedAt(),
		)

		mock.ExpectQuery("SELECT id, name, email, password_hash, phone, age, provider_id, is_2fa_enabled, is_email_verified, last_login_at, last_login_ip, last_login_country, last_login_city, created_at, updated_at FROM accounts WHERE email = \\$1 LIMIT 1").
			WithArgs(email).
			WillReturnRows(rows)

		repo := PGAccountsRepository{
			Db: db,
		}

		account := repo.FindByEmail(email)

		assert.NotNil(account)
		assert.Equal(email, account.GetEmail())

		err = mock.ExpectationsWereMet()
		assert.NoError(err)
	})

	t.Run("ErrorNotNoRows", func(t *testing.T) {
		mock.ExpectQuery("SELECT id, name, email, password_hash, phone, age, provider_id, is_2fa_enabled, is_email_verified, last_login_at, last_login_ip, last_login_country, last_login_city, created_at, updated_at FROM accounts WHERE email = \\$1 LIMIT 1").
			WithArgs(email).
			WillReturnError(errors.New("some other error"))

		repo := PGAccountsRepository{
			Db: db,
		}

		account := repo.FindByEmail(email)

		assert.Nil(account)

		err = mock.ExpectationsWereMet()
		assert.NoError(err)
	})
}

func TestPGAccountsRepository_Create(t *testing.T) {
	assert := assert.New(t)

	db, mock, err := sqlmock.New()
	assert.NoError(err)
	defer db.Close()

	hashedPass := "hashed-pass"
	phone := "999999999"
	providerId := ""

	repo := PGAccountsRepository{Db: db}

	accountData := entities.NewAccount("jhondoe", "jhondoe@gmail.com", hashedPass, phone, 23, providerId)

	t.Run("Success", func(t *testing.T) {
		mock.ExpectExec(`INSERT INTO accounts \(id, name, email, password_hash, phone, age, provider_id, created_at\) VALUES \(\$1, \$2, \$3, \$4, \$5, \$6, \$7, \$8\)`).
			WithArgs(
				accountData.GetID(),
				accountData.GetName(),
				accountData.GetEmail(),
				accountData.GetPass(),
				accountData.GetPhone(),
				accountData.GetAge(),
				accountData.GetProviderId(),
				accountData.GetCreatedAt(),
			).WillReturnResult(sqlmock.NewResult(1, 1))

		err = repo.Create(accountData)

		assert.NoError(err)

		err = mock.ExpectationsWereMet()
		assert.NoError(err)
	})

	t.Run("Error", func(t *testing.T) {
		mock.ExpectExec(`INSERT INTO accounts \(id, name, email, password_hash, phone, age, provider_id, created_at\) VALUES \(\$1, \$2, \$3, \$4, \$5, \$6, \$7, \$8\)`).
			WithArgs(
				accountData.GetID(),
				accountData.GetName(),
				accountData.GetEmail(),
				accountData.GetPass(),
				accountData.GetPhone(),
				accountData.GetAge(),
				accountData.GetProviderId(),
				accountData.GetCreatedAt(),
			).WillReturnError(errors.New("insert failed"))

		err = repo.Create(accountData)

		assert.Error(err)
		assert.Equal("insert failed", err.Error())

		err = mock.ExpectationsWereMet()
		assert.NoError(err)
	})
}

func TestPGAccountsRepository_Update(t *testing.T) {
	assert := assert.New(t)

	db, mock, err := sqlmock.New()
	assert.NoError(err)
	defer db.Close()

	hashedPass := "hashed-pass"
	phone := "999999999"
	providerId := ""

	accountData := entities.NewAccount("jhondoe", "jhondoe@gmail.com", hashedPass, phone, 23, providerId)

	repo := PGAccountsRepository{Db: db}

	t.Run("Success", func(t *testing.T) {
		mock.ExpectExec(`UPDATE accounts SET name = \$1, email = \$2, password_hash = \$3, phone_number = \$4, is_2fa_enabled = \$5, is_email_verified = \$6, last_login_at = \$7, last_login_ip = \$8, last_login_country = \$9, last_login_city = \$10, updated_at = \$11 WHERE id = \$12`).WithArgs(
			accountData.GetName(),
			accountData.GetEmail(),
			accountData.GetPass(),
			accountData.GetPhone(),
			accountData.GetIs2FaEnabled(),
			accountData.GetIsEmailVerified(),
			accountData.GetLastLoginAt(),
			accountData.GetLastLoginIp(),
			accountData.GetLastLoginCountry(),
			accountData.GetLastLoginCity(),
			accountData.GetUpdatedAt(),
			accountData.GetID(),
		).WillReturnResult(sqlmock.NewResult(1, 1))

		err = repo.Update(accountData)

		assert.NoError(err)

		err = mock.ExpectationsWereMet()
		assert.NoError(err)
	})

	t.Run("Error", func(t *testing.T) {
		mock.ExpectExec(`UPDATE accounts SET name = \$1, email = \$2, password_hash = \$3, phone_number = \$4, is_2fa_enabled = \$5, is_email_verified = \$6, last_login_at = \$7, last_login_ip = \$8, last_login_country = \$9, last_login_city = \$10, updated_at = \$11 WHERE id = \$12`).
			WithArgs(
				accountData.GetName(),
				accountData.GetEmail(),
				accountData.GetPass(),
				accountData.GetPhone(),
				accountData.GetIs2FaEnabled(),
				accountData.GetIsEmailVerified(),
				accountData.GetLastLoginAt(),
				accountData.GetLastLoginIp(),
				accountData.GetLastLoginCountry(),
				accountData.GetLastLoginCity(),
				accountData.GetUpdatedAt(),
				accountData.GetID(),
			).WillReturnError(errors.New("update failed"))

		err = repo.Update(accountData)

		assert.Error(err)
		assert.Equal("update failed", err.Error())

		err = mock.ExpectationsWereMet()
		assert.NoError(err)
	})
}
