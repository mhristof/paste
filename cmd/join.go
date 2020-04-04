package cmd

import (
	"fmt"

	"github.com/mhristof/paste/join"
	"github.com/spf13/cobra"
)

var (
	joinCmd = &cobra.Command{
		Use:   "join",
		Short: "Join the input lines",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			Verbose(cmd)
			fmt.Print(join.Join())
		},
	}
)

func init() {
	rootCmd.AddCommand(joinCmd)
}
