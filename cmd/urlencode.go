package cmd

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

var urlencodeCmd = &cobra.Command{
	Use:   "urlencode",
	Short: "URL encode a string",
	Run: func(cmd *cobra.Command, args []string) {
		data, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}

		fmt.Print(url.QueryEscape(string(data)))
	},
}

func init() {
	rootCmd.AddCommand(urlencodeCmd)
}
