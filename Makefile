.PHONY: install update test test_lib _test_lib test_cli _test_cli vet_lib _vet_lib vet_cli _vet_cli \
        fmt_check cli cli_doc cli_build fmt
.DEFAULT: help

STACK_NAME      = gogitlab
WIREMOCK_IMAGE  = ekino/wiremock:2.7.1
GO_IMAGE        = golang:1.10.3
MODD_VERSION    = 0.5
GO_PKG_SRC_PATH = "github.com/plouc/go-gitlab-client"
test_args      ?=

SHA1 = $(shell git rev-parse HEAD)
OS   = $(shell uname)

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
#  SETUP
#
#=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

setup: ##@setup Install all required components
	@echo "${YELLOW}Setting up project${RESET}"
	@${MAKE} clean
	@${MAKE} install_modd
	@${MAKE} up
	@${MAKE} install_go_deps
	@${MAKE} cli_all

install_modd: ##@setup Download modd watcher
	@mkdir -p ./bin
    ifeq ("$(wildcard ./bin/modd)", "")
		@echo "${YELLOW}Downloading modd watcher${RESET}"
        ifeq (${OS}, Darwin)
			@curl -L "https://github.com/cortesi/modd/releases/download/v${MODD_VERSION}/modd-${MODD_VERSION}-osx64.tgz" | tar xvf - --strip-components=1 -C ./bin
        endif
        ifeq (${OS}, Linux)
			@curl -L "https://github.com/cortesi/modd/releases/download/v${MODD_VERSION}/modd-${MODD_VERSION}-linux64.tgz" | tar xvzf - --strip-components=1 -C ./bin
        endif
		@echo ""
    endif

install_go_deps: ##@setup Install go dependencies
	@${MAKE} make_in_go TARGET=_install_go_deps

_install_go_deps:
	@echo "${YELLOW}Installing dependencies${RESET}"
	go get ${INSTALL_OPTS} gopkg.in/alecthomas/gometalinter.v2
	gometalinter.v2 --install
	go list -f '{{range .Imports}}{{.}} {{end}}' ./... | xargs go get ${INSTALL_OPTS}
	go list -f '{{range .TestImports}}{{.}} {{end}}' ./... | xargs go get ${INSTALL_OPTS}
	@echo "${GREEN}✔ successfully installed dependencies${RESET}\n"

update_go_deps: ##@setup Update dependencies
	@${MAKE} make_in_go TARGET=_update_go_deps

_update_go_deps:
	@${MAKE} _install_go_deps INSTALL_OPTS=-u

#=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
#
#  DOCKER
#
#=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

DOCKER_COMPOSE_VARS ?= export \
	GO_PKG_SRC_PATH=${GO_PKG_SRC_PATH} \
	GO_IMAGE=${GO_IMAGE} \
	WIREMOCK_IMAGE=${WIREMOCK_IMAGE};
DOCKER_COMPOSE_CMD ?= docker-compose -p ${STACK_NAME}
DOCKER_COMPOSE     := ${DOCKER_COMPOSE_VARS} ${DOCKER_COMPOSE_CMD}

up: ##@docker Launch docker compose stack in detached mode
	@${DOCKER_COMPOSE} up -t 0 --remove-orphans -d
	@echo ""

restart: ##@docker restart all services. To restart Quickly use QUICK=1
    ifdef QUICK
		@${DOCKER_COMPOSE} restart -t 0
    else
		@${DOCKER_COMPOSE} restart
    endif

restart_%: ##@docker restart % service. To restart Quickly use QUICK=1. eg. restart_wiremock
    ifdef QUICK
		@${DOCKER_COMPOSE} restart -t 0 ${*}
    else
		@${DOCKER_COMPOSE} restart ${*}
    endif

stop: ##@docker stop all services. To stop Quickly use QUICK=1
    ifdef QUICK
		@${DOCKER_COMPOSE} stop -t 0
    else
		@${DOCKER_COMPOSE} stop
    endif

stop_%: ##@docker stop % service. To stop Quickly use QUICK=1. eg. stop_wiremock
    ifdef QUICK
		@${DOCKER_COMPOSE} stop -t 0 ${*}
    else
		@${DOCKER_COMPOSE} stop ${*}
    endif

log_%: ##@docker Get stdout of running service. eg. log-wiremock
	@${DOCKER_COMPOSE} logs --tail=$${TAIL_LENGTH:-100} -f ${*}

log: ##@docker Print stdout of all services
	@${DOCKER_COMPOSE} logs -f --tail=$${TAIL_LENGTH:-50}

run_in_%: ##@docker Run a command inside % service
	@${MAKE} is_${*}_up
	@${DOCKER_COMPOSE} exec ${*} /bin/sh -c "${CMD}"

make_in_%: ##@docker Run a make command in % service
	@${MAKE} run_in_${*} CMD="${ENV} make ${TARGET}"

shell_in_%: ##@docker Get a shell in % service
	@${MAKE} is_${*}_up
	@${DOCKER_COMPOSE} exec ${*} /bin/sh

bash_in_%: ##@docker Get a shell (bash) in % service
	@${MAKE} is_${*}_up
	@${DOCKER_COMPOSE} exec ${*} /bin/bash

is_%_up: ##@docker make sure % service is up. eg. is_wiremock_up
	@/bin/sh -c ' \
	    IS_UP=`${DOCKER_COMPOSE} ps ${*} | grep Up`; \
	    if [ -z "$${IS_UP}" ]; then \
	        echo "${RED}✘ Service ${*} is down,${RESET}"; \
	        echo "${RED}  you should start stack through make up${RESET}"; \
	        exit 1; \
	    fi \
	'

wait_for_%: ##@docker make sure a service is up and reachable, eg. wait_for_wiremock WAIT_URL="http://crap:io" GREP_ITEM=ok
	@echo "${YELLOW}Waiting for ${WHITE}${*}${YELLOW} to be up and reachable${RESET}"
	@while true; do \
	    if curl -sS ${WAIT_URL} 2>/dev/null | grep -q ${GREP_ITEM}; then \
	        printf "\n"; \
	        echo "${GREEN}✔ ${*} is up!${RESET}" && exit 0; \
	    else \
	        printf "."; \
	    fi; \
	    sleep 1; \
	done;

ensure_wiremock_is_up: ##@docker wait for wiremock to be up
	@${MAKE} wait_for_wiremock WAIT_URL="http://wiremock:8080/__admin/mappings" GREP_ITEM="total"

clean: ##@docker Stop and remove docker compose stack
	@echo "${YELLOW}Cleaning docker-compose stack${RESET}"
	@${MAKE} stop QUICK=1
	@${DOCKER_COMPOSE} rm -f

status: ##@docker Docker compose services' status
	@echo "${YELLOW}Current docker compose services' status${RESET}"
	@${DOCKER_COMPOSE} ps

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
	@${MAKE} make_in_go TARGET=_test_lib

_test_lib:
	@echo "${YELLOW}Running lib tests${RESET}"
	@${MAKE} ensure_wiremock_is_up --no-print-directory
	@go test -v ${TESTS_OPTS} ./gitlab/.
	@echo "${GREEN}✔ Lib tests successfully passed${RESET}\n"

test_cli: ##@test Run CLI tests
	@${MAKE} make_in_go TARGET=_test_cli ENV="test_args='${test_args}'"

_test_cli:
	@echo "${YELLOW}Running CLI tests${RESET}"
	@${MAKE} ensure_wiremock_is_up --no-print-directory
	go test -v ./integration/. -args ${test_args}
	@echo "${GREEN}✔ CLI tests successfully passed${RESET}\n"

update_cli_snapshots: ##@test Update CLI test snapshots
	@${MAKE} make_in_go TARGET=_update_cli_snapshots

_update_cli_snapshots:
	@${MAKE} _test_cli test_args="-u all" --no-print-directory

vet_lib: ##@test Run vet on lib files
	@echo "${YELLOW}Running vet on lib${RESET}"
	@go vet ./gitlab/.
	@echo "${GREEN}✔ vet successfully passed for lib${RESET}\n"

vet_cli: ##@test Run vet on cli files
	@echo "${YELLOW}Running vet on cli${RESET}"
	@go vet ./cli/.
	@echo "${GREEN}✔ vet successfully passed for cli${RESET}\n"

fmt_check: ##@test Check formatting
	@${MAKE} make_in_go TARGET=_fmt_check

_fmt_check:
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
	@${MAKE} cli_doc
	@${MAKE} cli_build
	@${MAKE} cli_build_all

cli_doc: ##@cli Generate CLI documentation
	@${MAKE} make_in_go TARGET=_cli_doc

_cli_doc:
	@echo "${YELLOW}Generating CLI documentation${RESET}"
	@cd cli/doc && rm *.md && go run main.go
	@echo "${GREEN}✔ CLI documentation were successfully generated${RESET}\n"

cli_build: ##@cli Build CLI
	@${MAKE} make_in_go TARGET=_cli_build

_cli_build:
	@echo "${YELLOW}Building cli${RESET}"
	@cd cli && go build -o glc
	@echo "${GREEN}✔ successfully built CLI${RESET}\n"

cli_build_all: ##@cli Build CLI for various platforms
	@${MAKE} make_in_go TARGET=_cli_build_all

_cli_build_all:
	@echo "${YELLOW}Building CLI for various platforms${RESET}"
	cd cli && GOOS=darwin GOARCH=amd64 go build -o build/darwin-amd64-glc
	cd cli && GOOS=linux  GOARCH=amd64 go build -o build/linux-amd64-glc
	cd cli && GOOS=linux  GOARCH=386   go build -o build/linux-386-glc
	cd cli && GOOS=linux  GOARCH=arm   go build -o build/linux-arm-glc
	cd cli && GOOS=linux  GOARCH=arm64 go build -o build/linux-arm64-glc
	@echo "${GREEN}✔ successfully built CLI flavors${RESET}\n"

cli_checksums: ##@cli Generate checksums for CLI builds
	@${MAKE} make_in_go TARGET=_cli_checksums

_cli_checksums:
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
	@echo "${GREEN}✔ successfully generated CLI build checksums${RESET}\n"

#=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
#
#  MISC
#
#=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

fmt: ##@misc Format code
	@${MAKE} make_in_go TARGET=_fmt

_fmt:
	@echo "${YELLOW}Formatting code${RESET}"
	@gofmt -l -w -s .
	@go fix ./...
	@echo "${GREEN}✔ code was successfully formatted${RESET}\n"

dev: ##@misc Start watcher for development, auto run tests, fmt…
	@./bin/modd