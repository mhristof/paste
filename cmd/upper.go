package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	upperCmd = &cobra.Command{
		Use:   "upper",
		Short: "Convert to UPPER case",
		Run: func(cmd *cobra.Command, args []string) {
			Verbose(cmd)
			data, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				panic(err)
			}
			fmt.Println(strings.ToUpper(string(data)))
		},
	}
)

func init() {
	rootCmd.AddCommand(upperCmd)
}
