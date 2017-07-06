[![Build Status](https://travis-ci.org/dim13/xxhash.svg?branch=master)](https://travis-ci.org/dim13/xxhash)
[![GoDoc](https://godoc.org/github.com/dim13/xxhash?status.svg)](https://godoc.org/github.com/dim13/xxhash)

# xxhash is a non-cryptographic hash algorithm

This package implements [xxHash](https://cyan4973.github.io/xxHash/) in Go

```
BenchmarkXXHash32-8      1000000              1298 ns/op             416 B/op         52 allocs/op
BenchmarkXXHash64-8      2000000               721 ns/op             224 B/op         28 allocs/op
```
