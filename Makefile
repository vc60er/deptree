#This file is used for development cycle.
all: build

.PHONY: build
build:
	go build -v -o output/deptree

.PHONY: test
test:
	go test -v ./...

.PHONY: test_example
test_example:
	go test -v -run Example ./...

# Test, generate and show coverage in browser
.PHONY: test_cover
test_cover:
	mkdir -p output ; \
	go test -count=1 -v -run Test ./... -coverprofile=output/coverage.txt ; \
	go tool cover -html=output/coverage.txt ; \

.PHONY: clean
clean:
	@-rm -rf output/
	@-rm -rf .cache/