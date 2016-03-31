package main

import (
	"log"

	"github.com/juju/errors"
	"gitlab.altiplaconsulting.net/ernesto/rls/cmd"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(errors.ErrorStack(err))
	}
}

func run() error {
	if err := cmd.RootCmd.Execute(); err != nil {
		return errors.Trace(err)
	}

	return nil
}
