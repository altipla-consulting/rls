package cobracmd

import (
	"log"

	"git.altiplaconsulting.net/altipla/rls/rls"

	"github.com/juju/errors"
	"github.com/spf13/cobra"
)

var patchCmd = &cobra.Command{
	Use:   "patch",
	Short: "Release a new patch version",
	Run: func(cmd *cobra.Command, args []string) {
		if err := rls.Release(rls.Patch); err != nil {
			log.Fatal(errors.ErrorStack(err))
		}
	},
}

func init() {
	RootCmd.AddCommand(patchCmd)
}
