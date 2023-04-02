#This file is used for development cycle.
all: build

.PHONY: build
build:
	go build -o output/deptree -v ./cmd

.PHONY: clean
clean:
	@-rm -rf output/
	@-rm -rf .cache/