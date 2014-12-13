package main

import (
	goflags "github.com/jessevdk/go-flags"
	"log"
	"os"
)

// flags
var flags struct {
	Verbose bool `short:"v" long:"verbose" description:"Print verbose (debug) information" default:"false"`
}

// initFlags parses the given flags.
func initFlags() {
	args, err := goflags.Parse(&flags)
	if err != nil {
		os.Exit(1)
	}
	if len(args) > 0 {
		log.Fatalln("Unknown arguments:", args)
	}
}
