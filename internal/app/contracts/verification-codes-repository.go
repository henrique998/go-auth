package contracts

import "github.com/henrique998/go-auth/internal/app/entities"

type VerificationCodesRepository interface {
	FindByValue(val string) entities.VerificationCode
	Create(vc entities.VerificationCode) error
	Delete(codeId string) error
}
