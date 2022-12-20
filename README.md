# wdp

A Web Development Proxy with live reload capabilities.

## Installation

```bash
$ go install github.com/sho7a/wdp@latest
```

## Usage

```bash
Usage:
  wdp [flags]

Flags:
  -h, --help      help for wdp
  -v, --version   version for wdp
```

### Example

```bash
$ wdp -l 3000 -p 8080 -w bin
```

It works great with tools like [air](https://github.com/cosmtrek/air) or [nodemon](https://www.npmjs.com/package/nodemon).

## Build

```bash
$ git clone github.com/sho7a/wdp.git
$ cd wdp
$ make
```

## License

[Apache License 2.0](LICENSE)