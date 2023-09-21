package object

import (
	sender "github.com/casdoor/go-sms-sender"
	"strings"
)

func getSmsClient(provider *AdminProvider) (sender.SmsClient, error) {
	var client sender.SmsClient
	var err error

	client, err = sender.NewSmsClient(provider.Type, provider.ClientId, provider.ClientSecret, provider.SignName, provider.TemplateCode, "")
	if err != nil {
		return nil, err
	}

	return client, nil
}

func SendSms(provider *AdminProvider, content string, phoneNumbers ...string) error {
	client, err := getSmsClient(provider)
	if err != nil {
		return err
	}

	for i, number := range phoneNumbers {
		phoneNumbers[i] = strings.TrimPrefix(number, "+86")
	}

	params := map[string]string{}

	params["code"] = content

	err = client.SendMessage(params, phoneNumbers...)
	return err
}
