# Makefile

SHELL = /bin/bash

PACKAGE = $(shell cat go.mod | head -n 1 | cut -d ' ' -f 2)

BINS ?= $(shell ls cmd)

build: cmd/$(BINS)

package: cmd/$(BINS)/$(BINS).zip

cmd/$(BINS):
	go build \
		-o "$@/$(shell basename $@)-$(shell go env GOOS)-$(shell go env GOARCH)" \
		"$(PACKAGE)/$@"
.PHONY:	cmd/$(BINS)

cmd/$(BINS)/$(BINS).zip:
	zip $@ cmd/$(shell basename -s .zip $@)/$(shell basename -s .zip $@)-linux-amd64
.PHONY: cmd/$(BINS)/$(BINS).zip
