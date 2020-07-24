package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	lowerCmd = &cobra.Command{
		Use:   "lower",
		Short: "Convert to lower case",
		Run: func(cmd *cobra.Command, args []string) {
			Verbose(cmd)
			data, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				panic(err)
			}
			fmt.Print(strings.ToLower(string(data)))
		},
	}
)

func init() {
	rootCmd.AddCommand(lowerCmd)
}
