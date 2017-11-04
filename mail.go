package main

import (
	//"github.com/rulzurlibrary/api/utils"
	"github.com/spf13/cobra"
)

var Mail = &cobra.Command{
	Use:   "mail",
	Short: "Send mail using rulz mail configuration.",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		// Connect to the server, authenticate, set the sender and recipient,
		// and send the email all in one step.
		return ctx.Smtp.NewMail("default contact").
			To("foo", "usseppyrisa-7817@yopmail.com").
			Subject("Try some unicode éééé").
			Body([]byte("And here some more ààà")).
			Send()
	},
}

func init() {
	root.AddCommand(Mail)
}
