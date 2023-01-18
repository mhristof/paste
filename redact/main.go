package redact

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/tidwall/gjson"
	"github.com/trustelem/zxcvbn"
)

func Redact() string {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	return convert(data)
}

func convert(data []byte, whitelist []string) string {
	dataS := string(data)
	if gjson.Valid(dataS) {
		parsed := gjson.ParseBytes(data)
		encodeKeys(parsed)
	} else {
		text(dataS)
	}

	for secret, reduction := range redacted {
		dataS = strings.ReplaceAll(dataS, secret, reduction)
	}

	return dataS
}

func text(data string) {
	for _, word := range strings.FieldsFunc(data, split) {
		if len(word) > 100 {
			redacted[word] = redact(word)

			continue
		}

		res := zxcvbn.PasswordStrength(word, []string{})

		if res.CalcTime > 0.001 || res.Score > 1 {
			redacted[word] = redact(word)
		}
	}
}

func split(r rune) bool {
	return r == ' ' || r == '.' || r == '\n' || r == '=' || r == '?' || r == ':' || r == ',' || r == '"' || r == '\''
}

var redacted = map[string]string{}

func redact(in string) string {
	return fmt.Sprintf("sha256:%x", sha256.Sum256([]byte(in)))
}

func encodeKeys(result gjson.Result) bool {
	result.ForEach(func(key, value gjson.Result) bool {
		if value.IsObject() {
			return encodeKeys(value)
		}

		if len(value.String()) > 36 {
			redacted[strings.ReplaceAll(value.String(), "\n", `\n`)] = redact(value.String())

			return true
		}

		res := zxcvbn.PasswordStrength(value.String(), []string{})

		if res.CalcTime > 0.001 || res.Score > 1 {
			redacted[value.String()] = redact(value.String())

			return true
		}

		return true
	})

	return true
}
