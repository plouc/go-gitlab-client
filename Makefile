
STACK_NAME      = gogitlab
WIREMOCK_IMAGE  = ekino/wiremock:2.7.1
GO_IMAGE        = golang:1.10.3
MODD_VERSION    = 0.5
GO_PKG_SRC_PATH = "github.com/plouc/go-gitlab-client"

OS   = $(shell uname)

GIT_SHA  = $(shell git rev-parse HEAD)
GIT_REF ?= $(shell git symbolic-ref -q --short HEAD)

#=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
#
#  HELP
#
#=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

.DEFAULT: help

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
#  SETUP
#
#=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

install: ##@setup Install go dependencies
	@echo "${YELLOW}Installing dependencies${RESET}"
	go get ${install_flags} gopkg.in/alecthomas/gometalinter.v2
	gometalinter.v2 --install
	go list -f '{{range .Imports}}{{.}} {{end}}' ./... | xargs go get ${install_flags}
	go list -f '{{range .TestImports}}{{.}} {{end}}' ./... | xargs go get ${install_flags}
	@echo "${GREEN}✔ successfully installed dependencies${RESET}\n"

update: ##@setup Update dependencies
	@${MAKE} _install_go_deps install_flags=-u

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
	@echo "${GREEN}✔ well done!${RESET}\n"

test_lib: ##@test Run lib tests
	@echo "${YELLOW}Running lib tests${RESET}"
	@go test -v ${TESTS_OPTS} ./gitlab/.
	@echo "${GREEN}✔ Lib tests successfully passed${RESET}\n"

test_cli: ##@test Run CLI tests
	@echo "${YELLOW}Running CLI tests${RESET}"
	go test -v ./cli/cmd/. -args ${test_args}
	@echo "${GREEN}✔ CLI tests successfully passed${RESET}\n"

update_cli_snapshots: ##@test Update CLI test snapshots
	@${MAKE} test_cli test_args="-u all" --no-print-directory

vet_lib: ##@test Run vet on lib files
	@echo "${YELLOW}Running vet on lib${RESET}"
	@go vet ./gitlab/.
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

cli: ##@cli Run CLI from docker. eg. make cli CMD="ls groups"
	@${MAKE} run_in_go CMD="cd cli && ./glc ${CMD}"

cli_all: ##@cli Run all CLI steps
	@${MAKE} cli_build
	@${MAKE} cli_build_all

cli_build: ##@cli Build CLI
	@echo "${YELLOW}Building cli${RESET}"
	@cd cli && go build -o glc
	@echo "${GREEN}✔ successfully built CLI${RESET}\n"

cli_build_all: ##@cli Build CLI for various platforms
	@echo "${YELLOW}Building CLI for various platforms${RESET}"
	cd cli && GOOS=darwin GOARCH=amd64 go build -o build/darwin-amd64-glc
	cd cli && GOOS=linux  GOARCH=amd64 go build -o build/linux-amd64-glc
	cd cli && GOOS=linux  GOARCH=386   go build -o build/linux-386-glc
	cd cli && GOOS=linux  GOARCH=arm   go build -o build/linux-arm-glc
	cd cli && GOOS=linux  GOARCH=arm64 go build -o build/linux-arm64-glc
	@echo "${GREEN}✔ successfully built CLI flavors${RESET}\n"

cli_checksums: ##@cli Generate checksums for CLI builds
	@echo "${YELLOW}Generating CLI build checksums${RESET}"
	@rm -f cli/build/checksums.txt

    # for OSX users where you have md5 instead of md5sum
    ifeq (${OS}, Darwin)
        # md5 output has the following format:
        #
        # MD5 (darwin-amd64-glc) = 8eb317789e5d08e1c800cc469c20325a
        #
        # that's why sed and awk are used to cleanup
		@cd cli/build && ls . | grep glc \
            | xargs md5 \
            | awk '{ printf("%s\n%s\n\n", $$2, $$4) }' \
            | sed 's/[()]//g' \
            >> checksums.txt
    else
        # md5sum output has the following format:
        #
        # 8eb317789e5d08e1c800cc469c20325a darwin-amd64-glc
        #
        # that's why awk is used to cleanup
		@cd cli/build && ls . | grep glc \
            | xargs md5sum \
            | awk '{ printf("%s\n%s\n\n", $$2, $$1) }' \
            >> checksums.txt
    endif
	@echo "${GREEN}✔ successfully generated CLI build checksums to ${WHITE}cli/build/checksums.txt${RESET}\n"

#=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
#
#  MISC
#
#=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

fmt: ##@misc Format code
	@echo "${YELLOW}Formatting code${RESET}"
	@gofmt -l -w -s .
	@go fix ./...
	@echo "${GREEN}✔ code was successfully formatted${RESET}\n"

readme: ##@misc Generate README with CLI documentation
	@echo "${YELLOW}Generating ${WHITE}README.md${RESET}"
	@rm -f README.md
	@cat README.tpl.md >> README.md
	@go run cli/main.go doc >> README.md
	@echo "${GREEN}✔ successfully generated ${WHITE}README.md${RESET}\n"

.PHONY: install update \
        test test_lib test_cli update_cli_snapshots vet_cli vet_lib fmt_check \
        cli cli_all cli_build cli_build_all cli_checksums \
        fmt readme
