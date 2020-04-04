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
	out = removeIP(out)
	out = removeAWSHostname(out)
	out = removeUUID(out)
	out = removeARN(out)
	return out
}

func removeARN(lines string) string {
	padded := fmt.Sprintf(" %s ", lines)
	arn1 := regexp.MustCompile(`arn:\w*:\w*:\w*:\d*:\w*(\W)`)
	ret := arn1.ReplaceAllString(padded, "arn:partition:service:region:account-id:resource-id${1}")

	arn2 := regexp.MustCompile(`arn:\w*:\w*:\w*:\d*:\w*/\w*`)
	ret = arn2.ReplaceAllString(ret, "arn:partition:service:region:account-id:resource-type/resource-id")

	arn3 := regexp.MustCompile(`arn:\w*:\w*:\w*:\d*:\w*:\w*`)
	ret = arn3.ReplaceAllString(ret, "arn:partition:service:region:account-id:resource-type:resource-id")

	return strings.TrimSuffix(
		strings.TrimPrefix(ret, " "),
		" ",
	)
}

func removeUUID(lines string) string {
	uuidRE := regexp.MustCompile(`[0-9a-fA-F]{8}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{12}`)

	return uuidRE.ReplaceAllString(lines, "6433cfdf-43fa-4706-bdec-4d0b7872a68f")

}
func removeAWSHostname(lines string) string {
	hostname := regexp.MustCompile(`ip-\d{1,3}-\d{1,3}-\d{1,3}-(\d{1,3}).*\.internal`)

	res := hostname.ReplaceAllString(lines, "ip-123-123-123-${1}.region.service.internal")

	return res
}

func removeIP(lines string) string {
	ip := regexp.MustCompile(`\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}`)

	res := ip.ReplaceAllString(lines, "123.123.123.123")
	return res
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
