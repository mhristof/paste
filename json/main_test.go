package json

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
			name: "valid json input",
			in:   []byte(`{"this": "that"}`),
			out: heredoc.Doc(`
				{
				    "this": "that"
				}`),
		},
		{
			name: "valid json input with list",
			in:   []byte(`["this", "that"]`),
			out: heredoc.Doc(`
				[
				    "this",
				    "that"
				]`),
		},
		{
			name: "Invalid json input",
			in:   []byte(`this is not json`),
			out:  "this is not json",
		},
		{
			name: "Python output with single quotes",
			in:   []byte(`{'this': 'that'}`),
			out: heredoc.Doc(`
				{
				    "this": "that"
				}`),
		},
		{
			name: "Escaped json",
			in:   []byte(`{    \"this\": \"that\"}`),
			out: heredoc.Doc(`
				{
				    "this": "that"
				}`),
		},
		{
			name: "quoted escaped json",
			in:   []byte(`"{    \"this\": \"that\"}"`),
			out: heredoc.Doc(`
				{
				    "this": "that"
				}`),
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
