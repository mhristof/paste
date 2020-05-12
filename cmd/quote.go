package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	quoteCmd = &cobra.Command{
		Use:   "quote",
		Short: "Quote the input text with >",
		Run: func(cmd *cobra.Command, args []string) {
			Verbose(cmd)
			data, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				panic(err)
			}
			for _, line := range strings.Split(string(data), "\n") {
				fmt.Println(fmt.Sprintf("> %s", string(line)))
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(quoteCmd)
}
