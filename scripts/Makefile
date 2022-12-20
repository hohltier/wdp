.PHONY: all build

all: build

build:
ifeq ($(OS), Windows_NT)
	go build -o bin/wdp.exe
else
	go build -o bin/wdp
endif