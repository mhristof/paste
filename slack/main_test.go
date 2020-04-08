package slack

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
			name: "single line code",
			in:   []byte("this is a single line"),
			out:  "`this is a single line`",
		},
		{
			name: "multiline line code",
			in:   []byte("line one\nline two"),
			out:  "```\nline one\nline two\n```\n",
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
