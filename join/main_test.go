package join

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
			name: "simple two lines",
			in:   []byte("first\nsecond"),
			out:  "firstsecond",
		},
		{
			name: "multiline with backslases",
			in:   []byte("first \\ \nsecond"),
			out:  "first  second",
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
