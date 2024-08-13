package providers

import (
	"github.com/henrique998/go-auth/internal/app/contracts"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

type twilioTwoFactorAuthProvider struct {
	client *twilio.RestClient
}

func NewTwilioTwoFactorAuthProvider(client *twilio.RestClient) contracts.TwoFactorAuthProvider {
	return &twilioTwoFactorAuthProvider{client: client}
}

func (tp *twilioTwoFactorAuthProvider) Send(from, to, message string) error {
	params := &openapi.CreateMessageParams{}
	params.SetFrom(from)
	params.SetTo(to)
	params.SetBody(message)

	_, err := tp.client.Api.CreateMessage(params)
	if err != nil {
		return err
	}

	return nil
}
