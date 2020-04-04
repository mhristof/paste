package cmd

import (
	"fmt"

	"github.com/mhristof/paste/sha"
	"github.com/spf13/cobra"
)

var (
	shaCmd = &cobra.Command{
		Use:   "sha",
		Short: "Convert to sha256",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			Verbose(cmd)
			fmt.Print(sha.Sha())
		},
	}
)

func init() {
	rootCmd.AddCommand(shaCmd)
}
