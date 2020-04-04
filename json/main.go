package json

import (
	js "encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func Json() string {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	return convert(data)
}

func convert(lines []byte) string {
	out, err := bytesToJSON(lines)
	if err == nil {
		return out
	}

	// try for python objects which have single quotes
	out, err = bytesToJSON([]byte(strings.ReplaceAll(string(lines), "'", `"`)))
	if err == nil {
		return out
	}

	// escapped json
	out, err = bytesToJSON([]byte(strings.ReplaceAll(string(lines), `\"`, `"`)))
	if err == nil {
		return out
	}

	// quoted escapped json
	trimmed := strings.ReplaceAll(
		// Unquote the string
		strings.TrimSuffix(strings.TrimPrefix(string(lines), `"`), `"`),
		`\"`, `"`, // unescape the remaining "
	)
	out, err = bytesToJSON([]byte(trimmed))
	if err == nil {
		return out
	}

	return out
}

func bytesToJSON(lines []byte) (string, error) {
	var result map[string]interface{}

	js.Unmarshal(lines, &result)

	resultJSON, err := js.MarshalIndent(result, "", "    ")

	if err == nil && result != nil {
		return string(resultJSON), nil
	}

	var resultA []interface{}
	js.Unmarshal(lines, &resultA)

	resultJSON, err = js.MarshalIndent(resultA, "", "    ")
	if err == nil && resultA != nil {
		return string(resultJSON), nil
	}

	return string(lines), fmt.Errorf("Could not decode json text: %s", string(lines))
}
