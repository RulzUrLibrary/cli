package main

import (
	"github.com/spf13/cobra"
	"os"
	"path"
	"path/filepath"
)

const LIMIT = 50
const LIST_BOOKS = `SELECT isbn FROM books LIMIT $1 OFFSET $2`

var dry bool

func fileList(root string) (files []string, err error) {
	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.Mode().IsRegular() && info.Name()[0] != '.' {
			files = append(files, path)
		}
		return nil
	})
	return
}

func processThumbs(thumbs []string) ([]string, error) {
	for i := 0; ; i++ {
		var isbn string
		rows, err := ctx.Database.Query(LIST_BOOKS, LIMIT, i*LIMIT)

		if err != nil {
			return nil, err
		}

		for rows.Next() {
			if err := rows.Scan(&isbn); err != nil {
				return nil, err
			}
			thumb := path.Join(ctx.Configuration.Paths.Thumbs, isbn+".jpg")
			for i, t := range thumbs {
				if t == thumb {
					thumbs = append(thumbs[:i], thumbs[i+1:]...)
				}
			}
		}
		// if nothing has been scanned
		if isbn == "" {
			return thumbs, nil
		}
	}
}

var Clean = &cobra.Command{
	Use:   "clean",
	Short: "Clean some states from api runtime.",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		var thumbs []string

		if thumbs, err = fileList(ctx.Configuration.Paths.Thumbs); err != nil {
			return
		}

		if thumbs, err = processThumbs(thumbs); err != nil {
			return
		}

		ctx.Logger.Info("removing: ", thumbs)
		if !dry {
			for _, thumb := range thumbs {
				if err = os.Remove(thumb); err != nil {
					return
				}
			}
		}
		return nil
	},
}

func init() {
	dryMsg := "Dry run, log files removed but do not perform removal"
	Clean.PersistentFlags().BoolVar(&dry, "dry-run", false, dryMsg)

	root.AddCommand(Clean)
}
