# FIXME: add  option to strip debug symbols
testfinder: testfinder.go
	go build -ldflags '-s' testfinder.go

.PHONY: build
build: testfinder

.PHONY: test
test:
	go test
