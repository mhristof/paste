package cmd

import (
	"strings"

	"github.com/mhristof/paste/alfred"
	"github.com/spf13/cobra"
)

var (
	alfredCmd = &cobra.Command{
		Use:   "alfred",
		Short: "Print out the commands available in aflred format",
		Run: func(cmd *cobra.Command, args []string) {
			Verbose(cmd)
			var opts alfred.ScriptFilter

			for _, c := range rootCmd.Commands() {
				if c.Use == "alfred" || strings.HasPrefix(c.Use, "help") {
					continue
				}
				opts.Add(c.Use, c.Short)
			}
			opts.Print()
		},
	}
)

func init() {
	rootCmd.AddCommand(alfredCmd)
}
