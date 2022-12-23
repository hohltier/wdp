.PHONY: all build test test-build test-proxy

all: build

build:
ifeq ($(OS), Windows_NT)
	go build -o bin/wdp.exe
else
	go build -o bin/wdp
endif

test:
	$(MAKE) -j test-build test-proxy

test-build:
ifeq ($(OS), Windows_NT)
	air --build.cmd "go build -o bin/test.exe test/main.go"  --build.bin bin\test.exe
else
	air --build.cmd "go build -o bin/test test/main.go" --build.bin ./bin/test
endif

test-proxy:
ifeq ($(OS), Windows_NT)
	go build -o bin/wdp.exe
	bin\wdp.exe -l 3000 -p 8080 -w bin
else
	go build -o bin/wdp
	./bin/wdp -l 3000 -p 8080 -w bin
endif