BUILD_SHA=$(shell git rev-parse HEAD)
ALL_PACKAGES=$(shell go list ./... | grep -v /vendor/ | grep -v /scripts)

ifeq "$(UNAME)" "Darwin"
    BUILD_FLAGS=-ldflags="-s -X main.Build=$(BUILD_SHA)"
else
    BUILD_FLAGS=-ldflags="-X main.Build=$(BUILD_SHA)"
endif


CMD=github.com/glendc/go-gitignore/gitignore

build:
	go build $(BUILD_FLAGS) $(CMD)

BIN=./bin

NAME=gitignore

bin:
	mkdir -p $(BIN) 2> /dev/null

release: release-linux-386 release-linux-arm release-linux-amd64 release-darwin-386 release-darwin-amd64 release-windows-386 release-windows-amd64

release-linux-386: bin
	env GOARCH=386 GOOS=linux go build $(BUILD_FLAGS) \
		-o $(BIN)/$(NAME)-linux-386 $(CMD)

release-linux-arm: bin
	env GOARCH=arm GOOS=linux go build $(BUILD_FLAGS) \
		-o $(BIN)/$(NAME)-linux-arm $(CMD)

release-linux-amd64: bin
	env GOARCH=amd64 GOOS=linux go build $(BUILD_FLAGS) \
		-o $(BIN)/$(NAME)-linux-amd64 $(CMD)

release-darwin-386: bin
	env GOARCH=386 GOOS=darwin go build $(BUILD_FLAGS) \
		-o $(BIN)/$(NAME)-darwin-386 $(CMD)

release-darwin-amd64: bin
	env GOARCH=amd64 GOOS=darwin go build $(BUILD_FLAGS) \
		-o $(BIN)/$(NAME)-darwin-amd64 $(CMD)

release-windows-386: bin
	env GOARCH=386 GOOS=windows go build $(BUILD_FLAGS) \
		-o $(BIN)/$(NAME)-windows-386 $(CMD)

release-windows-amd64: bin
	env GOARCH=amd64 GOOS=windows go build $(BUILD_FLAGS) \
		-o $(BIN)/$(NAME)-windows-amd64 $(CMD)

clean:
	rm -rf $(GOPATH)/bin/$(NAME)

install: clean
	go install -v $(BUILD_FLAGS) $(CMD)

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
