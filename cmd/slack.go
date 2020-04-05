package cmd

import (
	"github.com/mhristof/paste/slack"
	"github.com/spf13/cobra"
)

var (
	slackCmd = &cobra.Command{
		Use:   "slack",
		Short: "Quote text with one or more `",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			Verbose(cmd)
			slack.Slack()
		},
	}
)

func init() {
	rootCmd.AddCommand(slackCmd)
}
