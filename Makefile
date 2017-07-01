PROJECT = github.com/nsknfd/go

ifneq ($(MAKECMDGOALS), clean)
SRC_FILE = $(shell find $(MAKECMDGOALS) -name "*.go")
endif

.PHONY: $(MAKECMDGOALS)
$(MAKECMDGOALS): $(SRC_FILE)
ifneq ($(MAKECMDGOALS), clean)
	@echo "Building $@ ..."
	@go build -o ${@F}.exe $(PROJECT)/$@
else
	@rm -rf *.exe
endif