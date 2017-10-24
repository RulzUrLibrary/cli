package main

import (
	//"github.com/rulzurlibrary/api/utils"
	"github.com/spf13/cobra"

	"fmt"
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
		//to := []string{"usseppyrisa-7817@yopmail.com"}
		//msg := []byte("From: default account <contact@rulz.xyz>\r\n" +
		//"To: usseppyrisa-7817@yopmail.com\r\n" +
		//"Subject: discount Gophers!\r\n" +
		//"\r\n" +
		//"This is the email body.\r\n")

		from := mail.Address{"监控中心", "fledna@163.com"}
		to := mail.Address{"收件人", "maxime.vidori@gmail.com"}
		title := "当前时段统计报表"

		body := "报表内容一切正常"

		header := make(map[string]string)
		header["From"] = from.String()
		header["To"] = to.String()
		header["Subject"] = encodeRFC2047(title)
		header["MIME-Version"] = "1.0"
		header["Content-Type"] = "text/plain; charset=\"utf-8\""

		message := ""
		for k, v := range header {
			message += fmt.Sprintf("%s: %s\r\n", k, v)
		}
		message += "\r\n" + body

		// Connect to the server, authenticate, set the sender and recipient,
		// and send the email all in one step.
		return ctx.Smtp.SendMail(
			[]string{to.Address},
			message,
			//[]byte("This is the email body."),
		)
	},
}

func init() {
	root.AddCommand(Mail)
}
