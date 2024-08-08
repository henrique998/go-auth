package createaccountcontroller

import (
	"github.com/gofiber/fiber/v3"
	"github.com/henrique998/go-auth/internal/app/request"
)

func (cc *createAccountController) Handle(c fiber.Ctx) error {
	req := request.CreateAccountRequest{
		Name:  "jhon doe",
		Email: "jhondoe@gmail.com",
		Pass:  "123456",
		Phone: "999999999",
	}

	cc.uc.Execute(req)

	return nil
}
