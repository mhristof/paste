package awscreds

import (
	"testing"

	"github.com/MakeNowJust/heredoc"
	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	var cases = []struct {
		name string
		in   []byte
		out  string
	}{
		{
			name: "partial text from aws sts assume-role command",
			in: []byte(heredoc.Doc(`
				"SecretAccessKey": "secret",
        		"SessionToken": "token",
        		"AccessKeyId": "access"`,
			)),
			out: "export AWS_SECRET_ACCESS_KEY='secret' AWS_ACCESS_KEY_ID='access' AWS_SESSION_TOKEN='token'",
		},
		{
			name: "root creds from the console",
			in: []byte(heredoc.Doc(`
				Access Key ID:
				access
				Secret Access Key:
				secret
			`)),
			out: "export AWS_ACCESS_KEY_ID='access' AWS_SECRET_ACCESS_KEY='secret'",
		},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			t.Log(tt.name)
			assert.Equal(t, tt.out, convert(tt.in))
		})
	}
}
