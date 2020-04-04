package cmd

import (
	"fmt"

	"github.com/mhristof/paste/json"
	"github.com/spf13/cobra"
)

var (
	jsonCmd = &cobra.Command{
		Use:   "json",
		Short: "Convert to pretty json",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			Verbose(cmd)
			fmt.Print(json.Json())
		},
	}
)

func init() {
	rootCmd.AddCommand(jsonCmd)
}
