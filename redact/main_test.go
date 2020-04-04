package redact

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
			name: "AWS key at the end",
			in:   []byte("This is an AWS secret key AKIA12345678C1234567"),
			out:  "This is an AWS secret key AKIAIOSXFEXAMPLE4567",
		},
		{
			name: "AWS key at the midle",
			in:   []byte("This is an AWS secret key AKIA12345678C1234567."),
			out:  "This is an AWS secret key AKIAIOSXFEXAMPLE4567.",
		},
		{
			name: "AWS key at the start",
			in:   []byte("AKIA12345678C1234567 is an AWS key"),
			out:  "AKIAIOSXFEXAMPLE4567 is an AWS key",
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
