package mail

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/selftechio/pigeon/internal/common"
)

const (
	// fixme 28/04/2021: hardcoded sender email address
	sender  = "dev-pigeon-mailer@pigeon-mail.selftech.io"
	charSet = "UTF-8"
)

type Mailer interface {
	SendMail(recipients []string, subject string, body string) error
}

type mailer struct {
	client *ses.SES
}

func NewMailer() Mailer {
	return &mailer{
		client: ses.New(common.Session),
	}
}

func (m *mailer) SendMail(recipients []string, subject string, body string) error {
	if len(recipients) == 0 {
		return nil
	}

	recipients_ptr := make([]*string, len(recipients))
	for i, recipient := range recipients {
		recipients_ptr[i] = aws.String(recipient)
	}

	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: recipients_ptr,
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: nil,
				Text: &ses.Content{
					Charset: aws.String(charSet),
					Data:    aws.String(body),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(charSet),
				Data:    aws.String(subject),
			},
		},
		Source: aws.String(sender),
	}

	_, err := m.client.SendEmail(input)
	// fixme 26/04/2021: find the exact aws error
	return err
}
