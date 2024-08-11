package verify2facodeusecase

import (
	"net/http"
	"testing"
	"time"

	"github.com/henrique998/go-auth/internal/app/entities"
	"github.com/henrique998/go-auth/internal/app/request"
	"github.com/henrique998/go-auth/internal/infra/utils"
	"github.com/henrique998/go-auth/test/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestVerify2faCodeUseCase(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAccountsRepo := mocks.NewMockAccountsRepository(ctrl)
	mockVCRepo := mocks.NewMockVerificationCodesRepository(ctrl)

	sut := NewVerify2faCodeUseCase(
		mockAccountsRepo,
		mockVCRepo,
	)

	t.Run("It should not be able to complete 2fa flow if code not exists", func(t *testing.T) {
		req := request.Verify2faCodeRequest{
			Code:      "",
			AccountId: "account-id",
		}

		mockVCRepo.EXPECT().FindByValue(req.Code).Return(nil)

		err := sut.Execute(req)

		assert.NotNil(err)
		assert.Equal("verification code not found", err.GetMessage())
		assert.Equal(http.StatusNotFound, err.GetStatus())
	})

	t.Run("It should not be able to complete 2fa flow if the account ID does not belong to the logged in account", func(t *testing.T) {
		codeStr, _ := utils.GenerateCode(10)
		code := entities.NewVerificationCode(codeStr, "invalid-account-id", time.Now().Add(10*time.Minute))

		req := request.Verify2faCodeRequest{
			Code:      codeStr,
			AccountId: "account-id",
		}

		mockVCRepo.EXPECT().FindByValue(codeStr).Return(code)

		err := sut.Execute(req)

		assert.NotNil(err)
		assert.Equal("unauthorized action", err.GetMessage())
		assert.Equal(http.StatusUnauthorized, err.GetStatus())
	})

	t.Run("It should not be able to complete 2fa flow it code has already expired", func(t *testing.T) {
		codeStr, _ := utils.GenerateCode(10)
		accountId := "account-id"
		code := entities.NewVerificationCode(codeStr, accountId, time.Now().Add(-1*time.Hour))

		mockVCRepo.EXPECT().FindByValue(codeStr).Return(code)

		req := request.Verify2faCodeRequest{
			Code:      codeStr,
			AccountId: accountId,
		}

		err := sut.Execute(req)

		assert.NotNil(err)
		assert.Equal("verification code has expired", err.GetMessage())
		assert.Equal(http.StatusUnauthorized, err.GetStatus())
	})

	t.Run("It should not be able to complete 2fa flow if already active", func(t *testing.T) {
		accountId := "account-id"
		providerId := ""
		lastLoginAt := time.Now().Add(-10 * time.Minute)
		lastLoginIp := "0.0.0.0"
		lastLoginCountry := "br"
		lastLoginCity := "sp"
		updatedAt := time.Now().Add(-5 * time.Hour)
		account := entities.NewExistingAccount(
			accountId,
			"jhon doe",
			"jhondoe@gmail.com",
			"123456",
			"999999999",
			23,
			&providerId,
			true,
			false,
			&lastLoginAt,
			&lastLoginIp,
			&lastLoginCountry,
			&lastLoginCity,
			time.Now().Add(-10*(time.Hour*24*10)),
			&updatedAt,
		)
		codeStr, _ := utils.GenerateCode(10)
		code := entities.NewVerificationCode(codeStr, accountId, time.Now().Add(10*time.Hour))

		mockVCRepo.EXPECT().FindByValue(codeStr).Return(code)
		mockAccountsRepo.EXPECT().FindById(accountId).Return(account)

		req := request.Verify2faCodeRequest{
			Code:      codeStr,
			AccountId: accountId,
		}

		err := sut.Execute(req)

		assert.NotNil(err)
		assert.Equal("Two factor authentication already carried out", err.GetMessage())
		assert.Equal(http.StatusUnauthorized, err.GetStatus())
	})

	t.Run("It should be able to complete 2fa", func(t *testing.T) {
		accountId := "account-id"
		providerId := ""
		lastLoginAt := time.Now().Add(-10 * time.Minute)
		lastLoginIp := "0.0.0.0"
		lastLoginCountry := "br"
		lastLoginCity := "sp"
		updatedAt := time.Now().Add(-5 * time.Hour)
		account := entities.NewExistingAccount(
			accountId,
			"jhon doe",
			"jhondoe@gmail.com",
			"123456",
			"999999999",
			23,
			&providerId,
			false,
			false,
			&lastLoginAt,
			&lastLoginIp,
			&lastLoginCountry,
			&lastLoginCity,
			time.Now().Add(-10*(time.Hour*24*10)),
			&updatedAt,
		)
		codeStr, _ := utils.GenerateCode(10)
		code := entities.NewVerificationCode(codeStr, accountId, time.Now().Add(10*time.Hour))

		mockVCRepo.EXPECT().FindByValue(codeStr).Return(code)
		mockVCRepo.EXPECT().Delete(code.GetID()).Return(nil)
		mockAccountsRepo.EXPECT().FindById(accountId).Return(account)
		mockAccountsRepo.EXPECT().Update(gomock.Any()).Return(nil)

		req := request.Verify2faCodeRequest{
			Code:      codeStr,
			AccountId: accountId,
		}

		err := sut.Execute(req)

		assert.Nil(err)
	})
}
