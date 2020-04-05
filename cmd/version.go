package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var GitCommit = "dev"

var (
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print version information",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(fmt.Sprintf("%s %s", rootCmd.Use, GitCommit))
		},
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)
}
