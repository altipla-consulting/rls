package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "rls",
	Short: "Rls releases new versions through Git tags in the repo using semver.",
}
