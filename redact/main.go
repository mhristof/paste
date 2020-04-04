package redact

import (
	"io/ioutil"
	"os"
	"regexp"
)

func Redact() string {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	return convert(data)
}

func convert(lines []byte) string {
	out := removeAWSCreds(string(lines))
	return out
}

func removeAWSCreds(lines string) string {
	access := regexp.MustCompile(`(\W{0,})[A-Z0-9]{16}([A-Z0-9]{4})(\W|$)`)
	res := access.ReplaceAllString(lines, "${1}AKIAIOSXFEXAMPLE${2}${3}")

	secret := regexp.MustCompile(`([A-Za-z0-9/+=]{40})(\W|$)`)
	res = secret.ReplaceAllString(res, "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

	return res
}
