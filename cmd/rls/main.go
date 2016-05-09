package main

import (
	"log"

	"github.com/juju/errors"
	"git.altiplaconsulting.net/altipla/rls/cobracmd"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(errors.ErrorStack(err))
	}
}

func run() error {
	if err := cobracmd.RootCmd.Execute(); err != nil {
		return errors.Trace(err)
	}

	return nil
}
