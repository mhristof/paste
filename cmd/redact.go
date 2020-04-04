package cmd

import (
	"fmt"

	"github.com/mhristof/paste/redact"
	"github.com/spf13/cobra"
)

var (
	redactCmd = &cobra.Command{
		Use:   "redact",
		Short: "Remove sensitive information",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			Verbose(cmd)
			fmt.Println(redact.Redact())

		},
	}
)

func init() {
	rootCmd.AddCommand(redactCmd)
}
