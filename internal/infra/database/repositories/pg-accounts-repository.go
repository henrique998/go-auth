package repositories

import (
	"database/sql"
	"time"

	"github.com/henrique998/go-auth/internal/app/entities"
	"github.com/henrique998/go-auth/internal/configs/logger"
)

type PGAccountsRepository struct {
	Db *sql.DB
}

type AccountRecord struct {
	ID               string
	Name             string
	Email            string
	Pass             *string
	Phone            string
	Age              int8
	ProviderId       *string
	Is2faEnabled     bool
	IsEmailVerified  bool
	LastLoginAt      *time.Time
	LastLoginIp      *string
	LastLoginCountry *string
	LastLoginCity    *string
	CreatedAt        time.Time
	UpdatedAt        *time.Time
}

func (r *PGAccountsRepository) FindById(accountId string) entities.Account {
	var accountData AccountRecord

	query := "SELECT id, name, email, password_hash, phone, age, provider_id, is_2fa_enabled, is_email_verified, last_login_at, last_login_ip, last_login_country, last_login_city, created_at, updated_at FROM accounts WHERE id = $1 LIMIT 1"
	row := r.Db.QueryRow(query, accountId)

	err := row.Scan(
		&accountData.ID,
		&accountData.Name,
		&accountData.Email,
		&accountData.Pass,
		&accountData.Phone,
		&accountData.Age,
		&accountData.ProviderId,
		&accountData.Is2faEnabled,
		&accountData.IsEmailVerified,
		&accountData.LastLoginAt,
		&accountData.LastLoginIp,
		&accountData.LastLoginCountry,
		&accountData.LastLoginCity,
		&accountData.CreatedAt,
		&accountData.UpdatedAt,
	)
	if err != nil {
		if err != sql.ErrNoRows {
			logger.Error("Error trying to find account by email", err)
		}
		return nil
	}

	account := entities.NewExistingAccount(
		accountData.ID,
		accountData.Name,
		accountData.Email,
		*accountData.Pass,
		accountData.Phone,
		accountData.Age,
		accountData.ProviderId,
		accountData.Is2faEnabled,
		accountData.IsEmailVerified,
		accountData.LastLoginAt,
		accountData.LastLoginIp,
		accountData.LastLoginCountry,
		accountData.LastLoginCity,
		accountData.CreatedAt,
		accountData.UpdatedAt,
	)

	return account
}

func (r *PGAccountsRepository) FindByEmail(email string) entities.Account {
	var accountData AccountRecord

	query := "SELECT id, name, email, password_hash, phone, age, provider_id, is_2fa_enabled, is_email_verified, last_login_at, last_login_ip, last_login_country, last_login_city, created_at, updated_at FROM accounts WHERE email = $1 LIMIT 1"
	row := r.Db.QueryRow(query, email)

	err := row.Scan(
		&accountData.ID,
		&accountData.Name,
		&accountData.Email,
		&accountData.Pass,
		&accountData.Phone,
		&accountData.Age,
		&accountData.ProviderId,
		&accountData.Is2faEnabled,
		&accountData.IsEmailVerified,
		&accountData.LastLoginAt,
		&accountData.LastLoginIp,
		&accountData.LastLoginCountry,
		&accountData.LastLoginCity,
		&accountData.CreatedAt,
		&accountData.UpdatedAt,
	)
	if err != nil {
		if err != sql.ErrNoRows {
			logger.Error("Error trying to find account by email", err)
		}
		return nil
	}

	account := entities.NewExistingAccount(
		accountData.ID,
		accountData.Name,
		accountData.Email,
		*accountData.Pass,
		accountData.Phone,
		accountData.Age,
		accountData.ProviderId,
		accountData.Is2faEnabled,
		accountData.IsEmailVerified,
		accountData.LastLoginAt,
		accountData.LastLoginIp,
		accountData.LastLoginCountry,
		accountData.LastLoginCity,
		accountData.CreatedAt,
		accountData.UpdatedAt,
	)

	return account
}

func (r *PGAccountsRepository) Create(a entities.Account) error {
	query :=
		`INSERT INTO accounts (id, name, email, password_hash, phone, age, provider_id, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err := r.Db.Exec(query,
		a.GetID(),
		a.GetName(),
		a.GetEmail(),
		a.GetPass(),
		a.GetPhone(),
		a.GetAge(),
		a.GetProviderId(),
		a.GetCreatedAt(),
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *PGAccountsRepository) Update(a entities.Account) error {
	query := "UPDATE accounts SET name = $1, email = $2, password_hash = $3, phone_number = $4, is_2fa_enabled = $5, is_email_verified = $6, last_login_at = $7, last_login_ip = $8, last_login_country = $9, last_login_city = $10, updated_at = $11 WHERE id = $12"

	_, err := r.Db.Exec(
		query,
		a.GetName(),
		a.GetEmail(),
		a.GetPass(),
		a.GetPhone(),
		a.GetIs2FaEnabled(),
		a.GetIsEmailVerified(),
		a.GetLastLoginAt(),
		a.GetLastLoginIp(),
		a.GetLastLoginCountry(),
		a.GetLastLoginCity(),
		a.GetUpdatedAt(),
		a.GetID(),
	)
	if err != nil {
		return err
	}

	return nil
}
