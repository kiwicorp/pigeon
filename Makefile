# Makefile

# Override the shell used.
SHELL = /bin/bash

# Targets to be manipulated by this Makefile.
TARGETS ?= $(shell ls cmd)

# Version to be used for various recipes.
VERSION ?= dev

build:
	@ ./scripts/build.sh $(VERSION) $(TARGETS)
.PHONY: build

package:
	@ ./scripts/package.sh $(VERSION) $(TARGETS)
.PHONY: package
