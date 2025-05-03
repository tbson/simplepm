package main

import (
	"fmt"
	"src/client/emailclient"
	"src/common/ctype"
	"src/module/account/extsrv/email"
)

func main() {
	client, err := emailclient.NewClient()
	if err != nil {
		fmt.Println(err)
		return
	}

	emailRepo := email.New(client)

	to := "tbson87@gmail.com"
	subject := "Test email for new client"
	body := ctype.EmailBody{
		HmtlPath: "emails/sample.html",
		Data: ctype.Dict{
			"Param": "Hello from Go!",
		},
	}

	emailRepo.SendEmailAsync(to, subject, body)

	fmt.Println("Press Enter to exit...")
	fmt.Scanln()
}
