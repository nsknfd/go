PROJECT = github.com/nsknfd/go

ifneq ($(MAKECMDGOALS), clean)
SRC_FILE = $(shell find $(MAKECMDGOALS) -name "*.go")
endif

.PHONY: $(MAKECMDGOALS)
$(MAKECMDGOALS): $(SRC_FILE)
ifneq ($(MAKECMDGOALS), clean)
	@echo "Building ${@F} ..."
	@go build -o main.exe $(PROJECT)/${@F}
else
	@rm -rf *.exe
endif