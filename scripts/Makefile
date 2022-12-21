.PHONY: all build test

all: build

build:
ifeq ($(OS), Windows_NT)
	go build -o bin/wdp.exe
else
	go build -o bin/wdp
endif

test: build
ifeq ($(OS), Windows_NT)
	bin/wdp.exe -l 3000 -p 8080 -w bin
else
	./bin/wdp -l 3000 -p 8080 -w bin
endif