package main

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/rulzurlibrary/api/app"
	"github.com/rulzurlibrary/api/ext/auth"
	"github.com/rulzurlibrary/api/ext/db"
	"github.com/rulzurlibrary/api/ext/scrapper"
	"github.com/spf13/cobra"
	"os"
)

type Context struct {
	Logger        *log.Logger
	Database      *db.DB
	Auth          *auth.Auth
	Scrapper      *scrapper.Scrapper
	Configuration app.Configuration
}

var (
	verbosity int
	debug     bool
	ctx       = &Context{}
	logLevel  = map[int]log.Lvl{0: log.OFF, 1: log.ERROR, 2: log.WARN, 3: log.INFO, 4: log.DEBUG}
	root      = &cobra.Command{
		Use:   "rulzctl",
		Short: "Central point of RulzurLibrary, deal with various components.",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			cmd.SilenceUsage = true // do not show usage if error come from RunE

			if verbosity > 4 || debug {
				verbosity = 4
			}

			ctx.Configuration = config.New()
			ctx.Logger = logger.New("rulzctl", logLevel[verbosity])
			ctx.Database = db.New(ctx.Logger, ctx.Configuration.Database)
			ctx.Auth = auth.New(ctx.Logger, ctx.Database)
			ctx.Scrapper = scrapper.New(ctx.Logger, ctx.Configuration.Paths.Thumbs)
		},
	}
)

func main() {
	debugMsg := "trigger debug logs, same as -vvvv, take precedence over verbose flag"
	verboseMsg := "verbose output, can be stacked to increase verbosity"

	root.PersistentFlags().BoolVar(&debug, "debug", false, debugMsg)
	root.PersistentFlags().CountVarP(&verbosity, "verbose", "v", verboseMsg)

	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
