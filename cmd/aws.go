package cmd

import (
	"fmt"

	"github.com/mhristof/paste/awscreds"
	"github.com/spf13/cobra"
)

var (
	awscredsCmd = &cobra.Command{
		Use:   "awscreds",
		Short: "Extract AWS creds from the text and coverte them to export stattements",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			Verbose(cmd)
			fmt.Print(awscreds.AWSCreds())
		},
	}
)

func init() {
	rootCmd.AddCommand(awscredsCmd)
}
