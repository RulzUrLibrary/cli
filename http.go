package main

import (
	"encoding/json"
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
		var notation utils.Notation

		for _, arg := range args {
			err = progress(fmt.Sprintf("scrapping book: %s  ", arg), 1, func(t *time.Ticker) error {
				defer t.Stop()

				book, err = ctx.Scrapper.Amazon(arg)
				if err != nil {
					return err
				}
				notation, err = ctx.Scrapper.SensCritique(book.TitleDisplay())
				if err != nil {
					return err
				}

				*book.Notations = append(*book.Notations, notation)
				return nil
			})
			if err != nil {
				return err
			}
			bytes, err := json.Marshal(book)
			if err != nil {
				return err
			}
			fmt.Printf("%s\n", bytes)
		}
		return nil
	},
}

func init() {
	root.AddCommand(Http)
}
