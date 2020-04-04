package base64

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	var cases = []struct {
		name string
		in   []byte
		out  string
	}{
		{
			name: "plain text",
			in:   []byte("VGhpcyBpcyBhIHN0cmluZw=="),
			out:  "This is a string",
		},
		{
			name: "plain text",
			in:   []byte("This is a string"),
			out:  "VGhpcyBpcyBhIHN0cmluZw==",
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
