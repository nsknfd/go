PROJECT = github.com/nsknfd/go

SRC_FILE = $(shell find src -name "*.go")

ALL = $(shell ls src)

all: $(ALL)
$(ALL): %: bin/%

bin/%: $(SRC_FILE)
	@echo "Building $@ ..."
	@go build -o $@ $(PROJECT)/src/${@F}

.PHONY: clean
	@rm -rf bin/