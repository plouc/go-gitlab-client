.PHONY: install update test test_lib _test_lib test_cli _test_cli vet_lib _vet_lib vet_cli _vet_cli \
        fmt_check cli cli_doc cli_completion cli_build fmt
.DEFAULT: help

STACK_NAME      = gogitlab
WIREMOCK_IMAGE  = ekino/wiremock:2.7.1
GO_IMAGE        = golang:1.10.3
MODD_VERSION    = 0.5
GO_PKG_SRC_PATH = "github.com/plouc/go-gitlab-client"
TESTS_OPTS     ?=

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
	@${MAKE} install_modd
	@${MAKE} up
	@${MAKE} install_go_deps
	@${MAKE} cli

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
	go list -f '{{range .Imports}}{{.}} {{end}}' ./... | xargs go get -v
	go list -f '{{range .TestImports}}{{.}} {{end}}' ./... | xargs go get -v
	@echo "${GREEN}✔ successfully installed dependencies${RESET}\n"

update_go_deps: ##@setup Update dependencies
	@${MAKE} make_in_go TARGET=_update_go_deps

_update_go_deps:
	@echo "${YELLOW}Updating dependencies${RESET}"
	go get -u all
	@echo "${GREEN}✔ successfully updated dependencies${RESET}\n"

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
    ifdef NO_DOCKER
		@${CMD}
    else
		@${MAKE} is_${*}_up
		@${DOCKER_COMPOSE} exec -T ${*} /bin/sh -c "${CMD}"
    endif

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
	@make ensure_wiremock_is_up --no-print-directory
	@go test ${TESTS_OPTS} ./gogitlab/.
	@echo "${GREEN}✔ Lib tests successfully passed${RESET}\n"

test_cli: ##@test Run CLI tests
	@${MAKE} make_in_go TARGET=_test_cli

_test_cli:
	@echo "${YELLOW}Running CLI tests${RESET}"
	@make ensure_wiremock_is_up --no-print-directory
	@go ${TESTS_OPTS} test ./integration/.
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

cli: ##@cli Run all cli steps
	@${MAKE} cli_doc
	@${MAKE} cli_completion
	@${MAKE} cli_build

cli_doc: ##@cli Generate cli documentation
	@${MAKE} make_in_go TARGET=_cli_doc

_cli_doc:
	@echo "${YELLOW}Generating cli documentation${RESET}"
	@cd cli/doc && rm *.md && go run main.go
	@echo "${GREEN}✔ cli documentation were successfully generated${RESET}\n"

cli_completion: ##@cli Generate cli bash completion
	@${MAKE} make_in_go TARGET=_cli_completion

_cli_completion:
	@echo "${YELLOW}Generating cli bash completion${RESET}"
	@cd cli/completion && go run main.go
	@echo "${GREEN}✔ cli bash completion were successfully generated${RESET}\n"

cli_build: ##@cli Build cli
	@${MAKE} make_in_go TARGET=_cli_build

_cli_build:
	@echo "${YELLOW}Building cli${RESET}"
	@cd cli && go build -o glc
	@echo "${GREEN}✔ successfully built CLI${RESET}\n"

cli_build_all: ##@cli Build cli for various platforms
	@${MAKE} make_in_go TARGET=_cli_build_all

_cli_build_all:
	@echo "${YELLOW}Building cli for various platforms${RESET}"
	cd cli && GOOS=darwin GOARCH=amd64 go build -o build/darwin-amd64-glc
	cd cli && GOOS=linux  GOARCH=amd64 go build -o build/linux-amd64-glc
	cd cli && GOOS=linux  GOARCH=386   go build -o build/linux-386-glc
	cd cli && GOOS=linux  GOARCH=arm   go build -o build/linux-arm-glc
	cd cli && GOOS=linux  GOARCH=arm64 go build -o build/linux-arm64-glc
	@echo "${GREEN}✔ successfully built CLI flavors${RESET}\n"

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

dev: ##@misc Start watcher for development, auto run tests, fmt…
	@./bin/modd