<div align="center">
    <a href="https://github.com/sho7a/wdp/blob/master/assets/figlet.txt"><img src="https://github.com/sho7a/wdp/raw/master/assets/figlet.png" alt="wdp" width="300" height="auto"></a>
    <h1>wdp</h1>
    <a href="https://github.com/sho7a/wdp/tags"><img alt="Version" src="https://img.shields.io/github/v/release/sho7a/wdp?label=Version"></a>
    <a href="https://github.com/sho7a/wdp/actions/workflows/build.yml"><img alt="Build" src="https://github.com/sho7a/wdp/actions/workflows/build.yml/badge.svg"></a>
    <a href="https://github.com/sho7a/wdp/blob/master/LICENSE"><img alt="License" src="https://img.shields.io/github/license/sho7a/wdp?label=License"></a>
</div>

## Features

- [x] Refresh browser
- [x] Detect recursive file changes
- [ ] Hot reload css

## Installation

```bash
$ go install github.com/sho7a/wdp@latest
```

## Usage

```bash
Usage:
  wdp [flags]

Flags:
  -h, --help           help for wdp
  -l, --listen int     listen port (default open port)
  -p, --port int       server port (default 80)
  -v, --version        version for wdp
  -w, --watch string   watch path (default ".")
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

[Apache License 2.0](https://github.com/sho7a/wdp/blob/master/LICENSE)