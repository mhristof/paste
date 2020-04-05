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

paste: $(shell find ./ -name '*.go')
	go build -o paste main.go

zip: paste
	zip -r paste.alfredworkflow info.plist paste 

help:           ## Show this help.
	@grep '.*:.*##' Makefile | grep -v grep  | sort | sed 's/:.*## /:/g' | column -t -s:
