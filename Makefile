APPNAME := $(shell basename $(CURDIR))

ifeq (,$(shell git describe --abbrev=0 --tags 2>/dev/null))
VERSION := v0.0.0
else
VERSION := $(shell git describe --abbrev=0 --tags)
endif

ifeq (,$(shell git rev-parse --short HEAD 2>/dev/null))
REVISION := unknown
else
REVISION := $(shell git rev-parse --short HEAD)
endif

GOOS   := linux
GOARCH := amd64

LDFLAGS_NAME     := -X "main.AppName=$(APPNAME)"
LDFLAGS_VERSION  := -X "main.Version=$(VERSION)"
LDFLAGS_REVISION := -X "main.Revision=$(REVISION)"
LDFLAGS          := -ldflags '-s -w $(LDFLAGS_NAME) $(LDFLAGS_VERSION) $(LDFLAGS_REVISION) -extldflags -static'

SRCS := $(shell find $(CURDIR) -type f -name '*.go')

.PHONY: all
all: $(APPNAME)

.PHONY: $(APPNAME)
$(APPNAME): $(CURDIR)/bin/$(APPNAME)
$(CURDIR)/bin/$(APPNAME): $(SRCS)
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build $(LDFLAGS) -o $@

$(CURDIR)/bin/$(APPNAME).zip: $(CURDIR)/bin/$(APPNAME)
	cd $(CURDIR)/bin && zip $@ $(APPNAME)

.PHONY: run
run: $(CURDIR)/bin/$(APPNAME)
	$(CURDIR)/bin/$(APPNAME)

.PHONY: install
install:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go install $(LDFLAGS)

.PHONY: test
test:
	go test -v

.PHONY: clean
clean:
	rm -rf $(CURDIR)/bin

.PHONY: init
init:
	go mod edit -module github.com/takumin/$(APPNAME)

.PHONY: patch-release
patch-release:
	git bump --patch

.PHONY: minor-release
minor-release:
	git bump --minor

.PHONY: major-release
major-release:
	git bump --major
