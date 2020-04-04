package cmd

import (
	"github.com/mhristof/paste/base64"
	"github.com/spf13/cobra"
)

var (
	base64Cmd = &cobra.Command{
		Use:   "base64",
		Short: "Convert from and to base64",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			Verbose(cmd)
			base64.Base64()

		},
	}
)

func init() {
	rootCmd.AddCommand(base64Cmd)
}
