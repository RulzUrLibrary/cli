package main

import (
	"github.com/rulzurlibrary/api/utils"
	"github.com/spf13/cobra"
)

var Captcha = &cobra.Command{
	Use:   "captcha",
	Short: "Try to fullfil the database with captcha blocked requests.",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		var book utils.Book
		var isbns []string

		if isbns, err = ctx.Database.CaptchaList(); err != nil {
			return err
		}
		for _, isbn := range isbns {
			ctx.Logger.Infof("processing isbn: %s", isbn)
			if book, err = ctx.Scrapper.Amazon(isbn); err != nil {
				return err
			}
			if err = ctx.Database.BookSave(&book); err != nil {
				return err
			}
			if _, err = ctx.Database.CaptchaRemove(isbn); err != nil {
				return err
			}
		}
		return nil
	},
}

func init() {
	root.AddCommand(Captcha)
}
