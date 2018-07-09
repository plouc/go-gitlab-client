SHA1=$(shell git rev-parse HEAD)

.PHONY: install update test test_lib test_cli vet_lib vet_cli fmt_check cli cli_doc cli_completion cli_build fmt

.DEFAULT: help

TEST_OPTS ?=

#=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
#
#  HELP
#
#=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

# COLORS
RED    = $(shell printf "\33[31m")
GREEN  = $(shell printf "\33[32m")
WHITE  = $(shell printf "\33[37m")
YELLOW = $(shell printf "\33[33m")
RESET  = $(shell printf "\33[0m")

# Add the following 'help' target to your Makefile
# And add help text after each target name starting with '\#\#'
# A category can be added with @category
HELP_SCRIPT = \
    %help; \
    while(<>) { push @{$$help{$$2 // 'options'}}, [$$1, $$3] if /^([a-zA-Z\-\%_]+)\s*:.*\#\#(?:@([a-zA-Z\-\%]+))?\s(.*)$$/ }; \
    print "usage: make [target]\n\n"; \
    for (sort keys %help) { \
    print "${WHITE}$$_:${RESET}\n"; \
    for (@{$$help{$$_}}) { \
    $$sep = " " x (32 - length $$_->[0]); \
    print "  ${YELLOW}$$_->[0]${RESET}$$sep${GREEN}$$_->[1]${RESET}\n"; \
    }; \
    print "\n"; }

help: ##prints help
	@perl -e '${HELP_SCRIPT}' ${MAKEFILE_LIST}

#=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
#
#  DEPS
#
#=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

install: ##@deps Install dependencies
	@echo "${YELLOW}Installing dependencies${RESET}"
	@go get github.com/stretchr/testify/assert
	@go list -f '{{range .Imports}}{{.}} {{end}}' ./... | xargs go get -v
	@go list -f '{{range .TestImports}}{{.}} {{end}}' ./... | xargs go get -v

update: ##@deps Update dependencies
	@echo "${YELLOW}Updating dependencies${RESET}"
	@go get -u all

#=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
#
#  TEST
#
#=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

test: ##@test Run all test steps
	@echo "${YELLOW}Running all tests${RESET}\n"
	@${MAKE} test_lib
	@${MAKE} test_cli
	@${MAKE} vet_lib
	@${MAKE} vet_cli
	@${MAKE} fmt_check
	@echo "${GREEN}✔ well done!${RESET}"

test_lib: ##@test Run lib tests
	@echo "${YELLOW}Running lib tests${RESET}"
	@go test -v ./gogitlab/.
	@echo "${GREEN}✔ Lib tests successfully passed${RESET}\n"

test_cli: ##@test Run CLI tests
	@echo "${YELLOW}Running CLI tests${RESET}"
	@go test -v ./integration/.
	@echo "${GREEN}✔ CLI tests successfully passed${RESET}\n"

vet_lib: ##@test Run vet on lib files
	@echo "${YELLOW}Running vet on lib${RESET}"
	@go vet ./gogitlab/.
	@echo "${GREEN}✔ vet successfully passed for lib${RESET}\n"

vet_cli: ##@test Run vet on cli files
	@echo "${YELLOW}Running vet on cli${RESET}"
	@go vet ./cli/.
	@echo "${GREEN}✔ vet successfully passed for cli${RESET}\n"

fmt_check: ##@test Check formatting
	@echo "${YELLOW}Checking formatting${RESET}"
	@exit `gofmt -l -s -e . | wc -l`
	@echo "${GREEN}✔ code was formatted as expected${RESET}\n"

#=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
#
#  CLI
#
#=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

cli: ##@cli Run all cli steps
	@${MAKE} cli_doc
	@${MAKE} cli_completion

cli_doc: ##@cli Generate cli documentation
	@echo "${YELLOW}Generating cli documentation${RESET}"
	@cd cli/doc && rm *.md && go run main.go
	@echo "${GREEN}✔ cli documentation were successfully generated${RESET}"

cli_completion: ##@cli Generate cli bash completion
	@echo "${YELLOW}Generating cli bash completion${RESET}"
	@cd cli/completion && go run main.go
	@echo "${GREEN}✔ cli bash completion were successfully generated${RESET}"

cli_build: ##@cli Build cli
	@echo "${YELLOW}Building cli${RESET}"
	@cd cli && go build -o glc
	@echo "${GREEN}✔ successfully built CLI${RESET}"

#=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
#
#  MISC
#
#=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

fmt: ##@misc Format code
	@echo "${YELLOW}Formatting code${RESET}"
	@gofmt -l -w -s .
	@go fix ./...
