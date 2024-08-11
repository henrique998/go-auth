package sendnewpassrequestusecase

import (
	"net/http"
	"testing"

	"github.com/henrique998/go-auth/internal/app/entities"
	"github.com/henrique998/go-auth/test/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestSendNewPassRequestUseCase(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAccountsRepo := mocks.NewMockAccountsRepository(ctrl)
	mockVCRepo := mocks.NewMockVerificationCodesRepository(ctrl)
	mockEmailProvider := mocks.NewMockEmailProvider(ctrl)

	usecase := NewSendNewPassRequestUseCase(
		mockAccountsRepo,
		mockVCRepo,
		mockEmailProvider,
	)

	t.Run("It should not be able to send new pass request if account not exists", func(t *testing.T) {
		email := "invalid-email"

		mockAccountsRepo.EXPECT().FindByEmail(email).Return(nil)

		err := usecase.Execute(email)

		assert.NotNil(err)
		assert.Equal("account does not exists", err.GetMessage())
		assert.Equal(http.StatusNotFound, err.GetStatus())
	})

	t.Run("It should be able to send new pass request", func(t *testing.T) {
		account := entities.NewAccount("jhon doe", "jhondoe@gmail.com", "123456", "999999999", 23, "")

		mockAccountsRepo.EXPECT().FindByEmail(account.GetEmail()).Return(account)
		mockVCRepo.EXPECT().Create(gomock.Any()).Return(nil)
		mockEmailProvider.EXPECT().SendMail(account.GetEmail(), "Password reset", gomock.Any())

		err := usecase.Execute(account.GetEmail())

		assert.Nil(err)
	})
}
