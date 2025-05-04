package email

import (
	"log"
	"time"

	"src/common/ctype"
	"src/common/setting"
	"src/util/errutil"
	"src/util/i18nmsg"
	"src/util/templateutil"

	"github.com/wneessen/go-mail"
)

type adapter struct {
	client *mail.Client
}

func New() adapter {
	emailClient, err := mail.NewClient(
		setting.EMAIL_HOST(),                             // SMTP server host
		mail.WithPort(setting.EMAIL_PORT()),              // SMTP server port
		mail.WithSMTPAuth(mail.SMTPAuthPlain),            // Auth mechanism
		mail.WithUsername(setting.EMAIL_HOST_USER()),     // Username
		mail.WithPassword(setting.EMAIL_HOST_PASSWORD()), // Password
		mail.WithTLSPolicy(mail.TLSMandatory),            // Force TLS
		mail.WithTimeout(30*time.Second),                 // Connection timeout
	)
	if err != nil {
		panic(err)
	}

	return adapter{
		client: emailClient,
	}
}

func (a adapter) SendEmail(to string, subject string, body ctype.EmailBody) error {
	message := mail.NewMsg()

	if err := message.From(setting.DEFAULT_EMAIL_FROM()); err != nil {
		return errutil.NewWithArgs(
			i18nmsg.FailedToSetFromAddress,
			ctype.Dict{
				"Value": setting.DEFAULT_EMAIL_FROM(),
			},
		)
	}

	if err := message.To(to); err != nil {
		return errutil.NewWithArgs(
			i18nmsg.FailedToSetToAddress,
			ctype.Dict{
				"Value": to,
			},
		)
	}

	message.Subject(subject)

	htmlBody, err := templateutil.GetHtmlString(body.HmtlPath, body.Data)
	if err != nil {
		return err
	}

	message.SetBodyString(mail.TypeTextHTML, htmlBody)

	if err := a.client.DialAndSend(message); err != nil {
		return errutil.New(i18nmsg.FailedToDeliverEmail)
	}

	return nil
}

func (a adapter) SendEmailAsync(to string, subject string, body ctype.EmailBody) {
	go func() {
		err := a.SendEmail(to, subject, body)
		if err != nil {
			log.Printf("Failed to send email to %s: %v", to, err)
		}
	}()
}
