package redact

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
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
	out = removeUser(out)
	out = removeInstanceID(out)
	return out
}

func removeInstanceID(lines string) string {
	oldID := regexp.MustCompile(`i-\w{5}(\w{3})`)
	res := oldID.ReplaceAllString(lines, "i-12345${1}")

	newID := regexp.MustCompile(`i-\w{14}(\w{3})`)
	res = newID.ReplaceAllString(res, "i-1234567890abcd${1}")

	return res
}

func removeUser(lines string) string {
	user := os.Getenv("USER")
	if user == "" {
		return lines
	}

	return strings.ReplaceAll(
		lines,
		user,
		"user",
	)
}

func removeAWSCreds(lines string) string {
	// instead of complicating the regex to take into consideration the tocken in
	// the start of the line and in the end, pad it with a known prefix and
	// remove it at the end
	padded := fmt.Sprintf("-%s-", lines)

	access := regexp.MustCompile(`(\W)[A-Z0-9]{16}([A-Z0-9]{4})(\W)`)
	res := access.ReplaceAllString(padded, "${1}AKIAIOSXFEXAMPLE${2}${3}")

	secret := regexp.MustCompile(`(\W)([A-Za-z0-9/+=]{35})([A-Za-z0-9/+=]{5})(\W)`)
	res = secret.ReplaceAllString(res, "${1}wJalrMI/K7MDENG/bPxRfiCYEXAMPLEKEY/${3}${4}")

	return strings.TrimSuffix(
		strings.TrimPrefix(res, "-"),
		"-",
	)
}
