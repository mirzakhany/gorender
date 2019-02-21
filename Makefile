NO_COLOR=\033[0m
OK_COLOR=\033[32;01m
ERROR_COLOR=\033[31;01m
WARN_COLOR=\033[33;01m

PKG_SRC := github.com/mirzakhany/gorender
PRG_NAME := gorender
export GOPATH=$(abspath $(HOME)/go)

.PHONY: all clean deps build

all: clean deps test build

deps:
	@echo "$(OK_COLOR)==> Installing dependencies$(NO_COLOR)"
#	@go get -u github.com/golang/dep/cmd/dep
#	@go get -u github.com/golang/lint/golint
	@dep ensure -v -vendor-only

build:
	@echo $(GOPATH)
	@echo "$(OK_COLOR)==> Building... $(NO_COLOR)"
	@/bin/sh -c "PKG_SRC=$(PKG_SRC) PROJECT_NAME=$(PRG_NAME) ./build/build.sh"

test: lint format vet
	@echo "$(OK_COLOR)==> Running tests$(NO_COLOR)"
	@go test -v -cover ./...

format:
	@echo "$(OK_COLOR)==> checking code formating with 'gofmt' tool$(NO_COLOR)"
	@gofmt -l -s cmd pkg | grep ".*\.go"; if [ "$$?" = "0" ]; then exit 1; fi

vet:
	@echo "$(OK_COLOR)==> checking code correctness with 'go vet' tool$(NO_COLOR)"
	@go vet ./...

lint: tools.golint
	@echo "$(OK_COLOR)==> checking code style with 'golint' tool$(NO_COLOR)"
	@go list ./... | xargs -n 1 golint -set_exit_status

clean:
	@echo "$(OK_COLOR)==> Cleaning project$(NO_COLOR)"
	@go clean
	@rm -rf dist $GOPATH/$(PKG_SRC)/dist

runbin:
	@echo "$(OK_COLOR)==> Running project$(NO_COLOR)"
	./dist/$(PRG_NAME)

run: clean build runbin


#---------------
#-- tools
#---------------

.PHONY: tools tools.dep tools.golint
tools: tools.dep tools.golint

tools.golint:
	@command -v golint >/dev/null ; if [ $$? -ne 0 ]; then \
		echo "--> installing golint"; \
		go get github.com/golang/lint/golint; \
	fi

tools.dep:
	@command -v dep >/dev/null ; if [ $$? -ne 0 ]; then \
		echo "--> installing dep"; \
		@go get -u github.com/golang/dep/cmd/dep; \
	fi
