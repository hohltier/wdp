.PHONY: all build test

all: build

build:
ifeq ($(OS), Windows_NT)
	go build -o bin/wdp.exe
else
	go build -o bin/wdp
endif

test:
	go run main.go