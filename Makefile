export SHELL := /usr/bin/env TZ=UTC bash

SOURCES := $(shell find . -name "*.go" | xargs)
PKGS := $(shell go list ./...)

UNITCOVEROUT := unitcoverage.out
COVERPKGS := $(shell echo "$(PKGS)" | sed 's, ,\,,')

INTTEST_DIR := inttest
INTCOVEROUT := $(INTTEST_DIR)/intcoverage
SUTBIN := $(INTTEST_DIR)/golang-test-lab

ifndef verbose
.SILENT:
endif

all: test

$(INTTEST_DIR) $(INTCOVEROUT):
	@echo "+ $@"
	mkdir -p $@

$(SUTBIN): $(SOURCES) | $(INTTEST_DIR)
	@echo "+ $@"
	(cd cmd; go build -cover -o $(CURDIR)/$@ -coverpkg=$(COVERPKGS) .)

show-coverpkgs:
	@echo "$(COVERPKGS)"

clean:
	@echo "+ $@"
	rm -f $(UNITCOVEROUT) $(SUTBIN)
	rm -Rf $(INTTEST_DIR)

# all tests
test: test-unit test-integration

test-unit $(UNITCOVEROUT):
	@echo "+ $@"
	go test -v --coverprofile=$(UNITCOVEROUT) $(PKGS)

test-unit-cov: $(UNITCOVEROUT)
	@echo "+ $@"
	go tool cover -func=$^

test-unit-cov-viz: $(UNITCOVEROUT)
	@echo "+ $@"
	go tool cover -html=$^

test-integration: $(SUTBIN) | $(INTCOVEROUT)
	@echo "+ $@"
	GOCOVERDIR=$(INTCOVEROUT) ./$(SUTBIN)

test-integration-cov: | $(INTCOVEROUT)
	@echo "+ $@"
	go tool covdata func -i=$(INTCOVEROUT)

.PHONY: all build clean show-coverpkgs test test-unit test-integration test-unit-cov test-integration-cov
