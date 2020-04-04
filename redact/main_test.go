package redact

import (
	"fmt"
	"os"
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
		{
			name: "Not an AWS key",
			in:   []byte("AAKIA12345678C1234567 is not an AWS key"),
			out:  "AAKIA12345678C1234567 is not an AWS key",
		},
		{
			name: "AWS secret at the end",
			in:   []byte("This is an AWS secret key wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY"),
			out:  "This is an AWS secret key wJalrMI/K7MDENG/bPxRfiCYEXAMPLEKEY/LEKEY",
		},
		{
			name: "AWS secret at the middle",
			in:   []byte("This is an AWS secret key wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY."),
			out:  "This is an AWS secret key wJalrMI/K7MDENG/bPxRfiCYEXAMPLEKEY/LEKEY.",
		},
		{
			name: "AWS secret at the start",
			in:   []byte("wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY is an AWS key"),
			out:  "wJalrMI/K7MDENG/bPxRfiCYEXAMPLEKEY/LEKEY is an AWS key",
		},
		{
			name: "not an AWS secret",
			in:   []byte("EwJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY is not a key"),
			out:  "EwJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY is not a key",
		},
		{
			name: "Multiple keys",
			in:   []byte("wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY and wJ11111tnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY are keys"),
			out:  "wJalrMI/K7MDENG/bPxRfiCYEXAMPLEKEY/LEKEY and wJalrMI/K7MDENG/bPxRfiCYEXAMPLEKEY/LEKEY are keys",
		},
		{
			name: "User name",
			in:   []byte(fmt.Sprintf("This is the user %s", os.Getenv("USER"))),
			out:  "This is the user user",
		},
		{
			name: "Old instance id",
			in:   []byte("This is an old instance id i-98765432"),
			out:  "This is an old instance id i-12345432",
		},
		{
			name: "New instance id",
			in:   []byte("This is a new instance id i-5c7b139248ad47f56"),
			out:  "This is a new instance id i-1234567890abcdf56",
		},
		{
			name: "IP addresses",
			in:   []byte("This is an ip address: 127.0.0.1"),
			out:  "This is an ip address: 123.123.123.123",
		},
		{
			name: "AWS hostname",
			in:   []byte("Hostname: ip-10-10-10-10.eu-west-1.compute.internal"),
			out:  "Hostname: ip-123-123-123-10.region.service.internal",
		},
		{
			name: "UUID",
			in:   []byte("UUID: 52bcbc59-e76f-4893-8beb-fc6f4923b81b"),
			out:  "UUID: 6433cfdf-43fa-4706-bdec-4d0b7872a68f",
		},
		{
			name: "ARN type 1",
			in:   []byte("arn:aws:iam::123456789012:user"),
			out:  "arn:partition:service:region:account-id:resource-id",
		},
		// {
		// 	name: "ARN type 2",
		// 	in:   []byte("arn:aws:iam::123456789012:user/foobar"),
		// 	out:  "arn:partition:service:region:account-id:resource-type/resource-id",
		// },
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
