package main

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/rulzurlibrary/api/app"
	"os"
	"path"
	"path/filepath"
	"time"
)

type _config struct{}
type _logger struct{}

var (
	config _config
	logger _logger
)

func abort(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (_ _config) New() app.Configuration {
	conf, err := app.ParseConfig()
	abort(err)

	// check if prod or dev env
	assets := path.Join("var", "lib", "rulzurlibrary")
	if _, err = os.Stat(assets); err != nil {
		assets, err = filepath.Abs(
			path.Join(path.Dir(os.Getenv(app.CONFIG_ENV)), "..", "assets"),
		)
		abort(err)
		_, err = os.Stat(assets)
		abort(err)
	}

	conf.View.Templates = path.Join(assets, "tplt")
	conf.Paths.Static = path.Join(assets, "static")
	conf.Paths.Thumbs = path.Join(assets, "thumbs")
	return conf
}

func (_ _logger) New(prefix string, level log.Lvl) *log.Logger {
	logger := log.New(prefix)
	logger.SetLevel(level)
	return logger
}

func progress(msg string, wait int, fn func(*time.Ticker) error) error {
	defer fmt.Print("\n")
	ticker := time.NewTicker(time.Second * time.Duration(wait))
	fmt.Print(msg)
	go func() {
		for _ = range ticker.C {
			fmt.Print(".")
		}
	}()
	return fn(ticker)
}
