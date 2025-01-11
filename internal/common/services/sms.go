package services

import (
	"fmt"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

type ISMSService interface {
	SendOTP(phoneNumber string, otp int64) error
	// SendOTPWithTemplate(phoneNumber string, otp int64, templateName string) error
}

type SMSService struct {
	client *twilio.RestClient
}

func NewSMSService(region string) *SMSService {

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: "",
		Password: "",
	})

	return &SMSService{
		client: client,
	}
}

func (s *SMSService) SendOTP(phoneNumber string, otp int64) error {
	message := fmt.Sprintf("Welcome to AstroKiran. Your OTP is: %d", otp)
	params := &openapi.CreateMessageParams{}
	params.SetTo(phoneNumber) // Recipient's phone number
	params.SetFrom("")        // Your Twilio phone number
	params.SetBody(message)

	// Send SMS
	resp, err := s.client.Api.CreateMessage(params)
	if err != nil {
		return fmt.Errorf("failed to send SMS: %v", err)
	}

	// Print response details
	fmt.Printf("SMS sent successfully: SID %s\n", *resp.Sid)
	return nil
}
