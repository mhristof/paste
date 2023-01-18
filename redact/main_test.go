package redact

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	cases := []struct {
		name string
		in   []byte
		out  string
	}{
		{
			name: "valid json with secrets",
			in:   []byte(`{"secret": "194B99BF-F227-4FF2-83F2-E9D925C1AAFF"}`),
			out:  `{"secret": "sha256:502de78865ad05f567de01e8c723375dcedbe9e37984ed3feceaa26f5f35a7b5"}`,
		},
		{
			name: "text with secrets",
			in:   []byte(`the secret is 194B99BF-F227-4FF2-83F2-E9D925C1AAFF`),
			out:  `the secret is sha256:502de78865ad05f567de01e8c723375dcedbe9e37984ed3feceaa26f5f35a7b5`,
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
