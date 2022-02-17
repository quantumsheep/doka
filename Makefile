VERSION ?= $$(git describe --tags 2>/dev/null || git rev-parse --short HEAD 2>/dev/null || echo "latest")
LEVEL ?= release
OUTPUT ?= doka$$(if [ "$${GOOS:-$$(go env GOOS)}" == "windows" ]; then echo '.exe'; else echo ''; fi)

GO_PACKAGE_PATH := github.com/quantumsheep/doka

ifeq ($(LEVEL),release)
GOLDFLAGS := -w -s
endif

build:
	go build -ldflags "$(GOLDFLAGS) -X '$(GO_PACKAGE_PATH)/cmd.Version=$(or $(strip $(VERSION)),latest)'" -o $(OUTPUT)

clean:
	rm -f doka

PREFIX ?= /usr/local

install: doka
	mkdir -p $(DESTDIR)$(PREFIX)/bin
	cp $< $(DESTDIR)$(PREFIX)/bin/doka

uninstall:
	rm -f $(DESTDIR)$(PREFIX)/bin/doka

default: build

.PHONY: clean install uninstall
