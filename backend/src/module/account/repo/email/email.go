package email

import (
	"src/common/ctype"
)

type Client interface {
	SendEmail(to string, subject string, body ctype.EmailBody) error
	SendEmailAsync(to string, subject string, body ctype.EmailBody)
}

type Repo struct {
	client Client
}

func New(client Client) Repo {
	return Repo{
		client: client,
	}
}

func (r Repo) SendEmail(to string, subject string, body ctype.EmailBody) error {
	return r.client.SendEmail(to, subject, body)
}

func (r Repo) SendEmailAsync(to string, subject string, body ctype.EmailBody) {
	r.client.SendEmailAsync(to, subject, body)
}
