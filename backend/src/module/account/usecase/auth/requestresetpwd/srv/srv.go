package srv

import (
	"src/common/ctype"
	"src/common/setting"
	"src/module/account/schema"
	"src/util/stringutil"
)

type userProvider interface {
	Retrieve(opts ctype.QueryOpts) (*schema.User, error)
	Update(opts ctype.QueryOpts, data ctype.Dict) (*schema.User, error)
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

func (srv srv) ResetPwdRequest(email string, tenantID uint) error {
	// Check user exists
	userOpts := ctype.QueryOpts{
		Filters: ctype.Dict{"Email": email, "TenantID": tenantID},
	}
	user, err := srv.userRepo.Retrieve(userOpts)
	if err != nil {
		return err
	}

	// Generate reset pwd token
	code := stringutil.GetRandomString(setting.OTP_LENGTH())

	// update user reset pwd token
	updateOpts := ctype.QueryOpts{Filters: ctype.Dict{"ID": user.ID}}
	updateData := ctype.Dict{
		"PwdResetToken": code,
	}
	_, err = srv.userRepo.Update(updateOpts, updateData)
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
