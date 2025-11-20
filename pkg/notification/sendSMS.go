package notification

import (
	"encoding/json"
	"fmt"

	"github.com/airlangga-hub/ecommerce-go/config"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)


type NotificationClient interface {
	SendSMS(phone, message string) error
}


type notificationClient struct {
	*config.AppConfig
}


func NewNotificationClient(cfg *config.AppConfig) NotificationClient {
	return &notificationClient{
		cfg,
	}
}


func (n *notificationClient) SendSMS(phone, message string) error {

	accountSid := n.TwilioAccountSid
	authToken := n.TwilioAuthToken

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	params := &twilioApi.CreateMessageParams{}
	params.SetTo(n.MyPhoneNumber)
	params.SetFrom(n.TwilioPhoneNumber)
	params.SetBody(message)

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println("Error sending SMS message: " + err.Error())
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Response: " + string(response))
	}

	return nil
}