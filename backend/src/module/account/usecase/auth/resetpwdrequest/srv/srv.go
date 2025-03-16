package srv

import (
	"src/common/ctype"
	"src/module/account/schema"
	"src/util/stringutil"
)

type userProvider interface {
	Retrieve(queryOptions ctype.QueryOptions) (*schema.User, error)
	Update(queryOptions ctype.QueryOptions, data ctype.Dict) (*schema.User, error)
}

type emailProvider interface {
	SendEmailAsync(to string, subject string, body ctype.EmailBody)
}

type srv struct {
	userRepo  userProvider
	emailRepo emailProvider
}

func New(userRepo userProvider, emailRepo emailProvider) srv {
	return srv{userRepo, emailRepo}
}

func (srv srv) RequestResetPwd(email string, tenantID uint) error {
	// Check user exists
	getUserOptions := ctype.QueryOptions{
		Filters: ctype.Dict{"Email": email, "TenantID": tenantID},
	}
	user, err := srv.userRepo.Retrieve(getUserOptions)
	if err != nil {
		return err
	}

	// Generate reset pwd token
	code := stringutil.GetRandomString(6)

	// update user reset pwd token
	updateOptions := ctype.QueryOptions{Filters: ctype.Dict{"ID": user.ID}}
	updateData := ctype.Dict{
		"PwdResetToken": code,
	}
	_, err = srv.userRepo.Update(updateOptions, updateData)
	if err != nil {
		return err
	}

	// Send email containing reset pwd token
	to := user.Email
	subject := "Reset Password"
	body := ctype.EmailBody{
		HmtlPath: "emails/reset-pwd.html",
		Data: ctype.Dict{
			"Name": user.FullName(),
			"Code": code,
		},
	}
	srv.emailRepo.SendEmailAsync(to, subject, body)
	return nil
}
