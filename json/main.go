package json

import (
	js "encoding/json"
	"io/ioutil"
	"os"
)

func Json() string {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	return convert(data)
}

func convert(lines []byte) string {
	var result map[string]interface{}

	js.Unmarshal(lines, &result)

	resultJSON, err := js.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	return string(resultJSON)
}
