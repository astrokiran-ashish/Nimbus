package services

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

type ISMSService interface {
	SendOTP(phoneNumber string, otp int64) error
	// SendOTPWithTemplate(phoneNumber string, otp int64, templateName string) error
}

type SMSService struct {
	snsClient *sns.SNS
}

func NewSMSService(region string) *SMSService {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
	}))
	return &SMSService{
		snsClient: sns.New(sess),
	}
}

func (s *SMSService) SendOTP(phoneNumber string, otp int64) error {
	message := fmt.Sprintf("Your OTP is: %d", otp)
	input := &sns.PublishInput{
		Message:     aws.String(message),
		PhoneNumber: aws.String(phoneNumber),
	}

	_, err := s.snsClient.Publish(input)
	return err
}
