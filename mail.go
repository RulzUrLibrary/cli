package main

import (
	//"github.com/rulzurlibrary/api/utils"
	"github.com/spf13/cobra"
	"net/mail"
	"strings"
)

func encodeRFC2047(String string) string {
	// use mail's rfc2047 to encode any string
	addr := mail.Address{String, ""}
	return strings.Trim(addr.String(), " <>")
}

var Mail = &cobra.Command{
	Use:   "mail",
	Short: "Send mail using rulz mail configuration.",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		// Connect to the server, authenticate, set the sender and recipient,
		// and send the email all in one step.
		return ctx.Smtp.NewMail("default contact").
			To("toto", "usseppyrisa-7817@yopmail.com").
			Subject("test test").
			Body("somebody").
			Send()
	},
}

func init() {
	root.AddCommand(Mail)
}
