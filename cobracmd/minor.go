package cobracmd

import (
	"log"

	"git.altiplaconsulting.net/altipla/rls/rls"

	"github.com/juju/errors"
	"github.com/spf13/cobra"
)

var minorCmd = &cobra.Command{
	Use:   "minor",
	Short: "Release a new minor version",
	Run: func(cmd *cobra.Command, args []string) {
		if err := rls.Release(rls.Minor); err != nil {
			log.Fatal(errors.ErrorStack(err))
		}
	},
}

func init() {
	RootCmd.AddCommand(minorCmd)
}
