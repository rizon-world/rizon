# Install Rizon Platform

### Build the source

#### Supported Systems

We currently supports the operating systems below.

* Ubuntu 18.04 or later
* MacOS 10.14 or later

#### Prerequisites

You should install the packages below before you build the source.

* [Golang](https://golang.org/doc/install) &gt;= 1.15
* make

#### Build

git clone the source and change directory.

```text
git clone https://github.com/rizon-world/rizon.git
cd rizon
```

Simply make it!

```text
make install
```

The built binary - _rizond_ - will be located in your `$GOBIN`.

Check the version of built binary.

```bash
$ rizond version
v0.1.0
```

#### License

RIZON PLATFORM is licensed under the [Apache License 2.0](https://github.com/rizon-world/rizon/blob/master/LICENSE).

