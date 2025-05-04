package main

import (
	"fmt"
	"src/adapter/email"
	"src/common/ctype"
)

func main() {
	emailAdapter := email.New()

	to := "tbson87@gmail.com"
	subject := "Test email for new client"
	body := ctype.EmailBody{
		HmtlPath: "emails/sample.html",
		Data: ctype.Dict{
			"Param": "Hello from Go!",
		},
	}

	emailAdapter.SendEmailAsync(to, subject, body)

	fmt.Println("Press Enter to exit...")
	fmt.Scanln()
}
