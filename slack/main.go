package slack

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func Slack() string {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	return convert(data)
}

func convert(lines []byte) string {
	multiline := len(strings.Split(string(lines), "\n")) > 1

	if multiline {
		return fmt.Sprintf("```\n%s\n```", string(lines))
	}

	return fmt.Sprintf("`%s`", string(lines))
}
