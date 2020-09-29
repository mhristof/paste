package awscreds

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func AWSCreds() string {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	return convert(data)
}

func convert(lines []byte) string {

	secretKey := findItem(string(lines), `"SecretAccessKey": "(?P<secret>.*)"`, "secret")
	accessKey := findItem(string(lines), `"AccessKeyId": "(?P<access>.*)"`, "access")
	sessionToken := findItem(string(lines), `"SessionToken": "(?P<token>.*)"`, "token")

	ret := "export"
	if secretKey != "" {
		ret = fmt.Sprintf("%s AWS_SECRET_ACCESS_KEY='%s'", ret, secretKey)
	}

	if accessKey != "" {
		ret = fmt.Sprintf("%s AWS_ACCESS_KEY_ID='%s'", ret, accessKey)
	}

	if sessionToken != "" {
		ret = fmt.Sprintf("%s AWS_SESSION_TOKEN='%s'", ret, sessionToken)
	}

	if ret == "export" {
		id := false
		key := false
		for _, line := range strings.Split(string(lines), "\n") {
			if id {
				ret = fmt.Sprintf("%s AWS_ACCESS_KEY_ID='%s'", ret, line)
			}

			if key {
				ret = fmt.Sprintf("%s AWS_SECRET_ACCESS_KEY='%s'", ret, line)
			}

			if line == "Access Key ID:" {
				id = true
			} else {
				id = false
			}

			if line == "Secret Access Key:" {
				key = true
			} else {
				key = false
			}
		}
	}

	return ret
}

func findItem(lines string, regex string, name string) string {
	secret := regexp.MustCompile(regex)
	for _, line := range strings.Split(string(lines), "\n") {
		res := secret.FindStringSubmatch(string(line))
		for i, thisName := range secret.SubexpNames() {
			if len(res) < i {
				break
			}
			if thisName == name {
				return res[i]
			}
		}
	}
	return ""
}
