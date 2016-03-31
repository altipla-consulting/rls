package cmd

import (
	"log"

	"gitlab.altiplaconsulting.net/ernesto/rls/rls"

	"github.com/juju/errors"
	"github.com/spf13/cobra"
)

var majorCmd = &cobra.Command{
	Use:   "major",
	Short: "Release a new major version",
	Run: func(cmd *cobra.Command, args []string) {
		if err := rls.Release(rls.Patch); err != nil {
			log.Fatal(errors.ErrorStack(err))
		}
	},
}

func init() {
	RootCmd.AddCommand(majorCmd)
}
