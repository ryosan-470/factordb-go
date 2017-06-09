# FactorDB CLI
[![CircleCI](https://img.shields.io/circleci/project/github/ryosan-470/factordb-go.svg?style=flat-square)](https://circleci.com/gh/ryosan-470/factordb-go)
[![AppVeyor](https://img.shields.io/appveyor/ci/ryosan-470/factordb-go.svg?style=flat-square)](https://ci.appveyor.com/project/ryosan-470/factordb-go/)
[![license](https://img.shields.io/github/license/ryosan-470/factordb-go.svg?style=flat-square)](https://github.com/ryosan-470/factordb-go/blob/master/LICENSE.md)
[![Codecov](https://img.shields.io/codecov/c/github/ryosan-470/factordb-go.svg?style=flat-square)](https://codecov.io/gh/ryosan-470/factordb-go)
[![Github All Releases](https://img.shields.io/github/downloads/ryosan-470/factordb-go/total.svg?style=flat-square)](https://github.com/ryosan-470/factordb-go/releases)

The [FactorDB](https://factordb.com) is the database to store known factorizations for any number. 
This tool can use on your command line.
Although I have already written [factordb-pycli](https://github.com/ryosan-470/factordb-pycli) with Python before, I want to use this tool on several platforms like on Windows machine and ARM Linux.

## Build & Installation (on macOS)
If you use macOS, you should follow the instruction.

### Requirements

* Go >= 1.8
* glide

```bash
$ brew install go glide
```

### Build

```bash
$ make deps
$ make
$ bin/factordb --help
NAME:
   factordb - The CLI for factordb.com

USAGE:
   factordb [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --json         Return response formated JSON
   --help, -h     show help
   --version, -v  print the version
```

After installation, you should copy a binary to your PATH.

## Windows
Download pre-build binary from [here](https://github.com/ryosan-470/factordb-go/releases).

| Platform | Status  |
|:---------|:-------:|
|Windows|[![AppVeyor](https://img.shields.io/appveyor/ci/ryosan-470/factordb-go.svg?style=flat-square)](https://ci.appveyor.com/project/ryosan-470/factordb-go/)|

## CLI
If you want to know the result of factorization of 16, you should type like this:

```bash
$ factordb 16
2 2 2 2
```

If you want to know more detail of result, you can get an answer of JSON format.

```bash
$ factordb --json 16
{"id": "https://factordb.com/?id=2", "status": "FF", "factors": [2, 2, 2, 2]}
```

# License
MIT
