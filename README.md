# crypt - A password hashing library for Go

crypt provides pure golang implementations of UNIX's crypt(3).

The goal of crypt is to bring a library of many common and popular password
hashing algorithms to Go and to provide a simple and consistent interface to
each of them. As every hashing method is implemented in pure Go, this library
should be as portable as Go itself.

All hashing methods come with a test suite which verifies their operation
against itself as well as the output of other password hashing implementations
to ensure compatibility with them.

I hope you find this library to be useful and easy to use!

*NOTE:* the package is a fork from 'github.com/GehirnInc/crypt'.

## Install

	go get github.com/tredoe/crypt@latest


## Documentation

The documentation is available on
[go.dev](https://pkg.go.dev/github.com/tredoe/crypt)

## License

The source files are distributed under the **BSD 2-Clause "Simplified" License**.
