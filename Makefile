# Makefile

# Override the shell used.
SHELL = /bin/bash

# Targets to be manipulated by this Makefile.
TARGETS ?= $(shell ls cmd)

# Version to be used for various recipes.
VERSION ?= dev

all: deploy

build:
	@ ./scripts/build.sh $(VERSION) $(TARGETS)
.PHONY: build

package: build
	@ ./scripts/package.sh $(VERSION) $(TARGETS)
.PHONY: package

deploy: package
	@ ./scripts/deploy.sh $(VERSION)
.PHONY: deploy

clean:
	@ ./scripts/clean.sh
.PHONY: clean
