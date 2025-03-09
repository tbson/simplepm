package emailclient

import (
	"log"
	"time"

	"src/common/ctype"
	"src/common/setting"
	"src/util/errutil"
	"src/util/localeutil"
	"src/util/templateutil"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/wneessen/go-mail"
)

type Client struct {
	client *mail.Client
}

func NewClient() (Client, error) {
	localizer := localeutil.Get()
	client, err := mail.NewClient(
		setting.EMAIL_HOST,                             // SMTP server host
		mail.WithPort(setting.EMAIL_PORT),              // SMTP server port
		mail.WithSMTPAuth(mail.SMTPAuthPlain),          // Auth mechanism
		mail.WithUsername(setting.EMAIL_HOST_USER),     // Username
		mail.WithPassword(setting.EMAIL_HOST_PASSWORD), // Password
		mail.WithTLSPolicy(mail.TLSMandatory),          // Force TLS
		mail.WithTimeout(30*time.Second),               // Connection timeout
	)
	if err != nil {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.FailedToCreateEmailClient,
		})
		return Client{}, errutil.New("", []string{msg})
	}

	return Client{
		client: client,
	}, nil
}

func (c Client) SendEmail(to string, subject string, body ctype.EmailBody) error {
	localizer := localeutil.Get()
	message := mail.NewMsg()

	if err := message.From(setting.DEFAULT_EMAIL_FROM); err != nil {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.FailedToSetFromAddress,
			TemplateData: ctype.Dict{
				"Value": setting.DEFAULT_EMAIL_FROM,
			},
		})
		return errutil.New("", []string{msg})
	}

	if err := message.To(to); err != nil {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.FailedToSetToAddress,
			TemplateData: ctype.Dict{
				"Value": to,
			},
		})
		return errutil.New("", []string{msg})
	}

	message.Subject(subject)

	htmlBody, err := templateutil.GetHtmlString(body.HmtlPath, body.Data)
	if err != nil {
		return err
	}

	message.SetBodyString(mail.TypeTextHTML, htmlBody)

	if err := c.client.DialAndSend(message); err != nil {
		msg := localizer.MustLocalize(&i18n.LocalizeConfig{
			DefaultMessage: localeutil.FailedToDeliverEmail,
		})
		return errutil.New("", []string{msg})
	}

	return nil
}

func (c Client) SendEmailAsync(to string, subject string, body ctype.EmailBody) {
	go func() {
		err := c.SendEmail(to, subject, body)
		if err != nil {
			log.Printf("Failed to send email to %s: %v", to, err)
		}
	}()
}
