package cmd

import (
	"fmt"
	"os"

	"github.com/mhristof/paste/log"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "paste",
	Short: "Manipulate the clipboard",
	Long:  `TODO: changeme`,
	Run: func(cmd *cobra.Command, args []string) {
		Verbose(cmd)
	},
}

// Verbose Increase verbosity
func Verbose(cmd *cobra.Command) {
	verbose, err := cmd.Flags().GetBool("verbose")
	if err != nil {
		log.Panic(err)
	}

	if verbose {
		log.SetLevel(log.DebugLevel)
	}
}
func init() {
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Increase verbosity")
}

// Execute The main function for the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
