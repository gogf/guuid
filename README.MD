# guuid

[![Go Doc](https://godoc.org/github.com/gogf/guuid?status.svg)](https://godoc.org/github.com/gogf/guuid)
[![License](https://img.shields.io/github/license/gogf/guuid.svg?style=flat)](https://github.com/gogf/guuid)

This package is a wrapper for most common used UUID package:
https://github.com/google/uuid

It supports the following UUID versions:
1. Version 1, based on timestamp and MAC address (RFC-4122)
2. Version 2, based on timestamp, MAC address and POSIX UID/GID (DCE 1.1)
3. Version 3, based on MD5 hashing of a named value (RFC-4122)
4. Version 4, based on random numbers (RFC-4122)
5. Version 5, based on SHA-1 hashing of a named value (RFC-4122)

## Installation

```shell
go get -u -v github.com/gogf/guuid@master
```

tip: when the `guuid` version is greater than `v1.0.0`, the dependent `GoFrame` version is `v2.1.1`

## License

The guuid is open-sourced software licensed under the [MIT license](./LICENSE).