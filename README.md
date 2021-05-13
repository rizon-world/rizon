# RIZON PLATFORM

[![license](https://img.shields.io/github/license/rizon-world/rizon.svg)](https://github.com/rizon-world/rizon/blob/master/LICENSE)

Welcome to the official RIZON PLATFORM repository.
RIZON is a decentralized network program which helps you to connect other blockchain and build your own network easily.

## Build the source

### Supported Systems

We currently supports the operating systems below.

* Ubuntu 18.04 or later
* MacOS 10.14 or later

### Prerequisites

You should install the packages below before you build the source.

* [Golang](https://golang.org/doc/install) >= 1.15
* make

### Build

git clone this source and change directory.

```sh
git clone https://github.com/rizon-world/rizon.git
cd rizon
```

Simply make it!

```sh
make install
```

The built binary - _rizond_ - will be located in your `$GOBIN`.

### License

RIZON PLATFORM is licensed under the [Apache License 2.0](https://github.com/rizon-world/rizon/blob/master/LICENSE).
