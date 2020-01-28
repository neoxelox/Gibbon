PACKAGES ?= ./...

GOTOOLS  ?= github.com/GeertJohan/fgt \
			golang.org/x/tools/cmd/goimports \
			github.com/kisielk/errcheck \
			golang.org/x/lint/golint \
			honnef.co/go/tools/cmd/staticcheck 

run:
	go run main.go

test:
	go test -race --count=1 $(PACKAGES)
.PHONY: test

check: test lint

lint: tools
	$(GOPATH)/bin/fgt go fmt $(PACKAGES)
	$(GOPATH)/bin/fgt go vet $(PACKAGES)
	$(GOPATH)/bin/fgt $(GOPATH)/bin/errcheck -ignore Close $(PACKAGES)
	echo $(PACKAGES) | xargs -L1 $(GOPATH)/bin/fgt $(GOPATH)/bin/golint
	$(GOPATH)/bin/staticcheck $(PACKAGES)
.SILENT: lint

tools:
	go get $(GOTOOLS)
.SILENT: tools