#
# vim:ft=make
#

SHELL := bash
.SHELLFLAGS := -eu -o pipefail -c
.ONESHELL:


GIT_COMMIT ?= $(shell git rev-list -1 HEAD)
BUILD_FLAGS := \
	-ldflags "-X github.com/mhristof/paste/cmd.GitCommit=$(GIT_COMMIT)"

fast-test:  ## Run fast tests
	go test ./... -tags fast

test:	## Run all tests
	go test ./...

paste: $(shell find ./ -name '*.go')
	go build $(BUILD_FLAGS) -o paste main.go

zip: paste
	zip -r paste.alfredworkflow info.plist paste 

v%:
	git tag v$*
	git push --tags

help:           ## Show this help.
	@grep '.*:.*##' Makefile | grep -v grep  | sort | sed 's/:.*## /:/g' | column -t -s:
