package main

import (
	"fmt"
	"github.com/rulzurlibrary/api/utils"
	"github.com/spf13/cobra"
	"time"
)

var Http = &cobra.Command{
	Use:   "http",
	Short: "Retrieve a book through its ISBN and display result.",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		var book utils.Book

		for _, arg := range args {
			progress(fmt.Sprintf("scrapping book: %s  ", arg), 1, func(t *time.Ticker) {
				book, err = ctx.Scrapper.Amazon(arg)
				t.Stop()
			})
			if err != nil {
				return err
			}
			fmt.Println(book)
		}
		return nil
	},
}

func init() {
	root.AddCommand(Http)
}
