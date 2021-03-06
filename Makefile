PROJECTNAME ?= desktop-helper
DESCRIPTION ?= desktop-helper - Desktop helper utility which can run scripts on new displays attached and more.
MAINTAINER  ?= Alexander Trost <galexrt@googlemail.com>
HOMEPAGE    ?= https://github.com/galexrt/desktop-helper

GO           := go
FPM          ?= fpm
PROMU        := $(GOPATH)/bin/promu
PREFIX       ?= $(shell pwd)
BIN_DIR      ?= $(PREFIX)/.build
TARBALL_DIR  ?= $(PREFIX)/.tarball
PACKAGE_DIR  ?= $(PREFIX)/.package
ARCH         ?= amd64
PACKAGE_ARCH ?= linux-amd64
VERSION      ?= $(shell cat VERSION)

pkgs = $(shell go list ./... | grep -v /vendor/ | grep -v /test/)

DOCKER_IMAGE_NAME ?= desktop-helper
DOCKER_IMAGE_TAG  ?= $(subst /,-,$(shell git rev-parse --abbrev-ref HEAD))

all: format style vet test build

build: promu
	@$(PROMU) build --prefix $(PREFIX)

crossbuild: promu
	@$(PROMU) crossbuild

docker:
	@echo ">> building docker image"
	@docker build -t "$(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG)" .

format:
	go fmt $(pkgs)

.PHONY: package
package-%: build
	mkdir -p -m0755 $(PACKAGE_DIR)/lib/systemd/system $(PACKAGE_DIR)/usr/bin
	mkdir -p $(PACKAGE_DIR)/etc/sysconfig
	cp .build/desktop-helper $(PACKAGE_DIR)/usr/bin
	cp systemd/desktop-helper.service $(PACKAGE_DIR)/lib/systemd/system
	cp systemd/sysconfig.desktop-helper $(PACKAGE_DIR)/etc/sysconfig/desktop-helper
	cd $(PACKAGE_DIR) && $(FPM) -s dir -t $(patsubst package-%, %, $@) \
	--deb-user root --deb-group root \
	--name $(PROJECTNAME) \
	--version $(VERSION) \
	--architecture $(PACKAGE_ARCH) \
	--description "$(DESCRIPTION)" \
	--maintainer "$(MAINTAINER)" \
	--url $(HOMEPAGE) \
	usr/ etc/

promu:
	@echo ">> fetching promu"
	@GOOS="$(shell uname -s | tr A-Z a-z)" \
	GOARCH="$(subst x86_64,amd64,$(patsubst i%86,386,$(shell uname -m)))" \
	$(GO) get -u github.com/prometheus/promu

style:
	@echo ">> checking code style"
	@! gofmt -d $(shell find . -path ./vendor -prune -o -name '*.go' -print) | grep '^'

tarball: promu
	@echo ">> building release tarball"
	@$(PROMU) tarball --prefix $(PREFIX) $(BIN_DIR)

test:
	@$(GO) test $(pkgs)

test-short:
	@echo ">> running short tests"
	@$(GO) test -short $(pkgs)

vet:
	@echo ">> vetting code"
	@$(GO) vet $(pkgs)

.PHONY: all build crossbuild docker format promu style tarball test vet
