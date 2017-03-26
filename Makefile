BUILD_SHA=$(shell git rev-parse HEAD)
ALL_PACKAGES=$(shell go list ./... | grep -v /vendor/ | grep -v /scripts)

ifeq "$(UNAME)" "Darwin"
    BUILD_FLAGS=-ldflags="-s -X main.Build=$(BUILD_SHA)"
else
    BUILD_FLAGS=-ldflags="-X main.Build=$(BUILD_SHA)"
endif

build:
	go build $(BUILD_FLAGS) \
		github.com/glendc/go-gitignore/gitignore

install:
	go install $(BUILD_FLAGS) \
		github.com/glendc/go-gitignore/gitignore

test:
ifeq "$(TRAVIS)" "true"
ifdef DARWIN
	sudo -E go test -v $(ALL_PACKAGES)
else
	go test $(BUILD_FLAGS) $(ALL_PACKAGES)
endif
else
	go test $(BUILD_FLAGS) $(ALL_PACKAGES)
endif
