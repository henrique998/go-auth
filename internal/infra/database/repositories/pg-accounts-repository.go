package repositories

import (
	"database/sql"

	"github.com/henrique998/go-auth-2/internal/app/entities"
)

type PGAccountsRepository struct {
	Db *sql.DB
}

func (r *PGAccountsRepository) FindById(accountId string) *entities.Account {
	return nil
}

func (r *PGAccountsRepository) FindByEmail(email string) *entities.Account {
	return nil
}

func (r *PGAccountsRepository) Create(a entities.Account) error {
	return nil
}

func (r *PGAccountsRepository) Update(a entities.Account) error {
	return nil
}
