package join

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func Join() string {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	return convert(data)
}

func convert(lines []byte) string {
	text := ""
	for _, line := range lines {
		sLine := string(line)
		text = fmt.Sprintf("%s%s", text, strings.TrimSuffix(sLine, "\n"))
	}
	text = strings.ReplaceAll(text, "\\", "")
	return text
}
