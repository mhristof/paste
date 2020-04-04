#
# vim:ft=make
#

SHELL := bash
.SHELLFLAGS := -eu -o pipefail -c
.ONESHELL:


fast-test:  ## Run fast tests
	go test ./... -tags fast

test:	## Run all tests
	go test ./...

help:           ## Show this help.
	@grep '.*:.*##' Makefile | grep -v grep  | sort | sed 's/:.*## /:/g' | column -t -s:
