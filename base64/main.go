package base64

import (
	"encoding/base64"
	"io/ioutil"
	"os"
)

func Base64() string {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	return convert(data)
}

func convert(lines []byte) string {
	out, err := base64.StdEncoding.DecodeString(string(lines))
	if err != nil {
		return base64.StdEncoding.EncodeToString(lines)
	}
	return string(out)
}
